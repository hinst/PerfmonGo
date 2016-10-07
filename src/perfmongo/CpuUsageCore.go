package perfmongo

type TCpuUsageCore struct {
	Total uint64
	Idle  uint64
}

type TCpuUsageCores []TCpuUsageCore

func (this *TCpuUsageCore) Substract(b TCpuUsageCore) {
	this.Total -= b.Total
	this.Idle -= b.Idle
}

func (this TCpuUsageCores) Substract(b TCpuUsageCores) TCpuUsageCores {
	var result = this.Clone()
	for i := range result {
		if i >= len(b) {
			break
		}
		result[i].Substract(b[i])
	}
	return result
}

func (this TCpuUsageCores) Clone() TCpuUsageCores {
	var result = make(TCpuUsageCores, len(this))
	copy(result, this)
	return result
}

func (this TCpuUsageCores) GetSum() TCpuUsageCore {
	var result TCpuUsageCore
	result.Total = 0
	result.Idle = 0
	for _, core := range this {
		result.Total += core.Total
		result.Idle += core.Idle
	}
	return result
}

func (this TCpuUsageCore) GetUtilization() float64 {
	return float64(this.Total-this.Idle) / float64(this.Total)
}
