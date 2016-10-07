package perfmongo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type TCpuUsageInfo struct {
	Moment time.Time
	Total  float64
	Cores  []float64
}

type TPlotlyLine struct {
	X []string  `json:"x"`
	Y []float64 `json:"y"`
}

type TPlotlyData struct {
	Series  []TPlotlyLine
	UnixNow int64
}

type TCpuUsageInfos []*TCpuUsageInfo

func (this TCpuUsageInfos) Clone() TCpuUsageInfos {
	var result = make(TCpuUsageInfos, len(this))
	copy(result, this)
	return result
}

func (this TCpuUsageInfos) GetLatest(duration time.Duration) TCpuUsageInfos {
	var result TCpuUsageInfos
	var now = time.Now()
	for _, item := range this {
		if now.Sub(item.Moment) < duration {
			result = append(result, item)
		}
	}
	return result
}

func (this TCpuUsageInfo) MomentToPlotlyString() string {
	var result = strconv.Itoa(this.Moment.Year()) + "-" +
		strconv.Itoa(int(this.Moment.Month())) + "-" +
		strconv.Itoa(this.Moment.Day()) + " " +
		strconv.Itoa(this.Moment.Hour()) + ":" +
		strconv.Itoa(this.Moment.Minute()) + ":" +
		strconv.Itoa(this.Moment.Second())
	return result
}

func (this TCpuUsageInfos) ToPlotlyJson() []byte {
	var line TPlotlyLine
	var dataObject TPlotlyData
	for _, item := range this {
		line.X = append(line.X, item.MomentToPlotlyString())
		line.Y = append(line.Y, item.Total*100)
	}
	dataObject.Series = append(dataObject.Series, line)
	dataObject.UnixNow = this.GetLastMoment().Unix()
	var data, jsonError = json.Marshal(dataObject)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	} else {
	}
	return data
}

func (this TCpuUsageInfos) GetLastMoment() time.Time {
	var result time.Time
	if len(this) > 0 {
		result = this[len(this)-1].Moment
	}
	return result
}

func (this TCpuUsageInfos) GetCountOfCores() int {
	var result = 0
	if len(this) > 0 {
		result = len(this[0].Cores)
	}
	return result
}

func (this TCpuUsageInfos) ToPlotlyCoresJson() []byte {
	var dataObject TPlotlyData
	var lines = make([]TPlotlyLine, this.GetCountOfCores())
	for _, item := range this {
		for coreIndex, core := range item.Cores {
			lines[coreIndex].X = append(lines[coreIndex].X, item.MomentToPlotlyString())
			lines[coreIndex].Y = append(lines[coreIndex].Y, core)
		}
	}
	dataObject.Series = lines
	dataObject.UnixNow = this.GetLastMoment().Unix()
	var data, _ = json.Marshal(dataObject)
	return data
}

func (this *TCpuUsageInfo) LoadFromDiff(diff TCpuUsageCores) {
	if len(diff) > 0 {
		this.Total = diff[0].GetUtilization()
		diff = diff[1:]
		this.Cores = make([]float64, len(diff))
		for i, item := range diff {
			this.Cores[i] = item.GetUtilization()
		}
	}
}
