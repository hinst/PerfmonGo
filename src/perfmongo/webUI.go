package perfmongo

//response http.ResponseWriter, request *http.Request

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type TWebUI struct {
	AppURL        string
	RequestHolder sync.WaitGroup
	Perfmon       TPerfmon
}

func (this *TWebUI) Start() {
	this.AppURL = "/PerfmonGo"
	this.AddRequestHandler("/page", this.HandlePageRequest)
	this.InstallFileHandler("third/web")
	go http.ListenAndServe(":9001", nil)
}

func (this *TWebUI) Stop() {
	this.RequestHolder.Wait()
}

func (this *TWebUI) ReadLayout() []byte {
	var content = GetCachedAsset("src/page/layout.html")
	return content
}

func (this *TWebUI) HandlePageRequest(response http.ResponseWriter, request *http.Request) {
	var pageName = request.URL.Query().Get("name")
	var page = GetCachedAsset("src/page/" + pageName + ".html")
	if page != nil {
		var layout = this.ReadLayout()
		var content = strings.Replace(string(layout), "$body", string(page), -1)
		content = strings.Replace(content, "$appURL", this.AppURL, -1)
		response.Write([]byte(content))
	} else {
		if false {
			fmt.Println("Error: could not load page '" + pageName + "'")
		}
	}
}

func (this *TWebUI) WrapRequestHandler(f func(response http.ResponseWriter, request *http.Request)) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		this.RequestHolder.Add(1)
		f(response, request)
		this.RequestHolder.Done()
	}
}

func (this *TWebUI) AddRequestHandler(subUrl string, f func(response http.ResponseWriter, request *http.Request)) {
	http.HandleFunc(this.AppURL+subUrl, this.WrapRequestHandler(f))
}

func (this *TWebUI) InstallFileHandler(folderName string) {
	var url = this.AppURL + "/" + folderName + "/"
	var directoryPath = AppDirectory + "/" + folderName
	var fileDirectory = http.Dir(directoryPath)
	var fileServerHandler = http.FileServer(fileDirectory)
	http.Handle(url, http.StripPrefix(url, fileServerHandler))
}
