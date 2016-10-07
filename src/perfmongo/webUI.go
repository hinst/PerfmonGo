package perfmongo

//response http.ResponseWriter, request *http.Request

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type TWebUI struct {
	AppURL          string
	RequestHolder   sync.WaitGroup
	Perfmon         *TPerfmon
	ThirdWebPrefix  []byte
	PagePrefix      []byte
	ThirdWebHandler fasthttp.RequestHandler
	PagePath        []byte
	PageSrcPrefix   []byte
	PageFiler       fasthttp.RequestHandler
	LatestCpuPath   []byte
}

func (this *TWebUI) Start() {
	this.AppURL = "/PerfmonGo"
	this.ThirdWebPrefix = []byte(this.AppURL + "/third/web")
	this.ThirdWebHandler = fasthttp.FSHandler(AppDirectory+"/third/web", 3)
	this.PagePath = []byte(this.AppURL + "/page")
	this.PageSrcPrefix = []byte(this.AppURL + "/src/page")
	this.PageFiler = fasthttp.FSHandler(AppDirectory+"/src/page", 3)
	this.LatestCpuPath = []byte(this.AppURL + "/latestCPU")
	go fasthttp.ListenAndServe(":9001", this.ProcessRequest)
}

func (this *TWebUI) ProcessRequest(ctx *fasthttp.RequestCtx) {
	this.RequestHolder.Add(1)
	defer this.RequestHolder.Done()
	var path = ctx.Path()
	switch {
	case bytes.HasPrefix(path, this.ThirdWebPrefix):
		this.ThirdWebHandler(ctx)
	case bytes.HasPrefix(path, this.PageSrcPrefix):
		this.PageFiler(ctx)
	case bytes.Equal(path, this.PagePath):
		this.HandlePageRequest(ctx)
	case bytes.HasPrefix(path, this.LatestCpuPath):
		this.GetLatest(ctx)
	default:
		ctx.SetBodyString("Not found")
	}
}

func (this *TWebUI) Stop() {
	this.RequestHolder.Wait()
}

func (this *TWebUI) ReadLayout() []byte {
	var content = GetCachedAsset("src/page/layout.html")
	return content
}

func (this *TWebUI) HandlePageRequest(ctx *fasthttp.RequestCtx) {
	var pageName = string(ctx.URI().QueryArgs().Peek("name"))
	var page = GetCachedAsset("src/page/" + pageName + ".html")
	if page != nil {
		var layout = this.ReadLayout()
		var content = strings.Replace(string(layout), "$body", string(page), -1)
		content = strings.Replace(content, "$appURL", this.AppURL, -1)
		ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
		ctx.SetBodyString(content)
	} else {
		if false {
			fmt.Println("Error: could not load page '" + pageName + "'")
		}
	}
}

func (this *TWebUI) GetLatest(ctx *fasthttp.RequestCtx) {
	var seconds = ctx.URI().QueryArgs().GetUintOrZero("seconds")
	var cores = ctx.URI().QueryArgs().GetUintOrZero("cores")
	if seconds > 0 {
		var duration = time.Second * time.Duration(seconds)
		this.Perfmon.DataLocker.RLock()
		defer this.Perfmon.DataLocker.RUnlock()
		var latestData = this.Perfmon.Data.GetLatest(duration)
		var plotlyData []byte
		if cores > 0 {
			plotlyData = latestData.ToPlotlyCoresJson()
		} else {
			plotlyData = latestData.ToPlotlyJson()
		}
		ctx.Response.SetBody(plotlyData)
	}
}
