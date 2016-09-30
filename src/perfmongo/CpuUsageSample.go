package perfmongo

type TCpuUsageSample struct {
	Idle  uint64
	Cores []uint64
}
