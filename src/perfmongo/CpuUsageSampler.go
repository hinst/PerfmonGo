package perfmongo

import (
	"io/ioutil"
	"strings"
)

type TCpuUsageSampler struct {
	Interval int
}

func (this *TCpuUsageSampler) ReadSystem() {
	var text, readFileResult = ioutil.ReadFile("/proc/stat")
	if readFileResult == nil {
		var lines = strings.Split(string(contents), "\n")
		for 
		var fields = strings.Fields(line)
	}
}
