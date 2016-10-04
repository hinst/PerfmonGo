package perfmongo

//response http.ResponseWriter, request *http.Request

import (
	"fmt"
	"io/ioutil"
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
	go http.ListenAndServe(":9001", nil)
}

func (this *TWebUI) Stop() {
	this.RequestHolder.Wait()
}

func (this *TWebUI) ReadLayout() []byte {
	var content, _ = ioutil.ReadFile(AppDirectory + "/src/page/layout.html")
	return content
}

func (this *TWebUI) CheckStrictPageName(pageName string) bool {
	var result = len(pageName) > 0
	if result {
		var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for _, c := range pageName {
			var acceptable = strings.ContainsRune(letters, c)
			if false == acceptable {
				result = false
				break
			}
		}
	}
	return result
}

func (this *TWebUI) HandlePageRequest(response http.ResponseWriter, request *http.Request) {
	var pageName = request.URL.Query().Get("name")
	if this.CheckStrictPageName(pageName) {
		var page, pageResult = GetAsset("src/page/" + pageName + ".html")
		if pageResult == nil {
			var layout = this.ReadLayout()
			var content = strings.Replace(string(layout), "$body", string(page), -1)
			content = strings.Replace(content, "$appURL", this.AppURL, -1)
			response.Write([]byte(content))
		} else {
			fmt.Println("Error: could not load page '" + pageName + "' " + pageResult.Error())
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
