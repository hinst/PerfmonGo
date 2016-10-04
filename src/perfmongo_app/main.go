package main

import (
	"fmt"
	"perfmongo"
	"sync"
)

var perfmon perfmongo.TPerfmon
var holder sync.WaitGroup

func main() {
	fmt.Println("STARTING...")
	holder.Add(1)
	perfmon.CountOfCores = 8
	perfmon.Start()
	perfmongo.InstallShutdownReceiver(func() {
		perfmon.Stop()
		holder.Done()
	})
	holder.Wait()
	fmt.Println("EXITING...")
}
