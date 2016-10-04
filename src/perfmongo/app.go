package perfmongo

type TApp struct {
	Perfmon TPerfmon
}

func (this *TApp) Start() {
	this.Perfmon.CountOfCores = 8
	this.Perfmon.Start()
}

func (this *TApp) Stop() {
	this.Perfmon.Stop()
}
