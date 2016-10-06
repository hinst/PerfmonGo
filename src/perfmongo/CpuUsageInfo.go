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
}

type TPlotlyData struct {
	X []string  `json:"x"`
	Y []float64 `json:"y"`
}

type TCpuUsageInfos []TCpuUsageInfo

func (this TCpuUsageInfos) Clone() TCpuUsageInfos {
	var result = make(TCpuUsageInfos, len(this))
	copy(result, this)
	return result
}

func (this TCpuUsageInfos) GetLatest(duration time.Duration) TCpuUsageInfos {
	fmt.Println(len(this))
	var result TCpuUsageInfos
	var now = time.Now()
	for _, item := range this {
		if now.Sub(item.Moment) < duration {
			fmt.Println("+")
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
		dataObject.X = append(dataObject.X, item.MomentToPlotlyString())
		dataObject.Y = append(dataObject.Y, item.Total)
	}
	var data, _ = json.Marshal(dataObject)
	return data
}
