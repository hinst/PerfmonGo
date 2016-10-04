package main

import (
	"fmt"
	"perfmongo"
	"sync"
)

var app perfmongo.TApp
var holder sync.WaitGroup

func main() {
	fmt.Println("STARTING...")
	perfmongo.GetAsset = Asset
	holder.Add(1)
	app.Start()
	perfmongo.InstallShutdownReceiver(func() {
		app.Stop()
		holder.Done()
	})
	holder.Wait()
	fmt.Println("EXITING...")
}
