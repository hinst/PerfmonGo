package perfmongo

import (
	"encoding/json"
	"strconv"
	"time"
)

type TCpuUsageInfo struct {
	Moment time.Time
	Total  float64
}

type TPlotlyData struct {
	x []string
	y []float64
}

type TCpuUsageInfos []TCpuUsageInfo

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
	var dataObject TPlotlyData
	for _, item := range this {
		dataObject.x = append(dataObject.x, item.MomentToPlotlyString())
		dataObject.y = append(dataObject.y, item.Total)
	}
	var data, _ = json.Marshal(dataObject)
	return data
}
