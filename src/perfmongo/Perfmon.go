package perfmongo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TPerfmon struct {
	CountOfCores int
	Interval     int
	Last         TCpuUsageCores
	thread       TCycleThread
	DataLocker   sync.RWMutex
	Data         []TCpuUsageInfo
}

func (this *TPerfmon) ReadSystem() TCpuUsageCores {
	var text, readFileResult = ioutil.ReadFile("/proc/stat")
	var result TCpuUsageCores = make(TCpuUsageCores, 8)
	var resultIndex = 0
	if readFileResult == nil {
		var lines = strings.Split(string(text), "\n")
		for _, line := range lines {
			var fields = strings.Fields(line)
			if len(fields) > 0 && len(fields[0]) >= 3 && fields[0][:3] == "cpu" {
				fields = fields[1:]
				for fieldIndex, field := range fields {
					var value, conversionResult = strconv.ParseUint(field, 10, 64)
					if conversionResult == nil {
						result[resultIndex].Total += value
					}
					if fieldIndex == 3 {
						result[resultIndex].Idle += value
					}
				}
				resultIndex++
			}
		}
	}
	return result
}

func (this *TPerfmon) Start() {
	this.Last = this.ReadSystem()
	this.thread.Interval = time.Second * 2
	this.thread.Function = this.execute
	this.thread.Start()
}

func (this *TPerfmon) Stop() {
	this.thread.Stop()
}

func (this *TPerfmon) execute() {
	var current = this.ReadSystem()
	var diff = current.Clone()
	diff.Substract(this.Last)
	this.Last = current
	var totalUtilization = diff[0].GetUtilization()
	if false {
		fmt.Printf("%v %v\n", diff[0].Total, diff[0].Idle)
		fmt.Println(strconv.FormatFloat(totalUtilization, 'f', 2, 64))
	}
	var data TCpuUsageInfo
	data.Moment = time.Now()
	data.Total = totalUtilization
	this.DataLocker.Lock()
	defer this.DataLocker.Unlock()
	this.AddData(data)
	this.ReduceData()
}

func (this *TPerfmon) AddData(info TCpuUsageInfo) {
	this.Data = append(this.Data, info)
}

func (this *TPerfmon) GetDataLengthLimit() int {
	return 10000
}

func (this *TPerfmon) GetDataLengthToCut() int {
	return 1000
}

func (this *TPerfmon) ReduceData() {
	if len(this.Data) > this.GetDataLengthLimit() {
		this.Data = this.Data[this.GetDataLengthToCut():]
	}
}
