package perfmongo

import (
	"os"
	"path/filepath"
)

var AppDirectory, _ = filepath.Abs(filepath.Dir(os.Args[0]))

type TApp struct {
	Perfmon TPerfmon
	Web     TWebUI
}

func (this *TApp) Start() {
	this.Perfmon.CountOfCores = 8
	this.Perfmon.Start()
	this.Web.Start()
}

func (this *TApp) Stop() {
	this.Web.Stop()
	this.Perfmon.Stop()
}
