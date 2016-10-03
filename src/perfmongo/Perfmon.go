package perfmongo

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type TPerfmon struct {
	Interval int
	thread   TCycleThread
}

func (this *TPerfmon) ReadSystem() []TCpuUsageCore {
	var text, readFileResult = ioutil.ReadFile("/proc/stat")
	var result []TCpuUsageCore
	if readFileResult == nil {
		var lines = strings.Split(string(text), "\n")
		for _, line := range lines {
			var fields = strings.Fields(line)
			if fields[0] == "cpu" {
				result = append(result, TCpuUsageCore{})
				fields = fields[1:]
				for fieldIndex, field := range fields {
					var value, conversionResult = strconv.ParseUint(field, 10, 64)
					if conversionResult == nil {
						result[len(result)-1].Total += value
					}
					if fieldIndex == 4 {
						result[len(result)-1].Idle += value
					}
				}
			}
		}
	}
	return result
}

func (this *TPerfmon) Start() {
	this.thread.Interval = time.Second * 2
	this.thread.Function = this.execute
	this.thread.Start()
}

func (this *TPerfmon) execute() {
}
