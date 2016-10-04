package perfmongo

import "time"

type TCpuUsageInfo struct {
	Moment time.Time
	Total  float64
}
