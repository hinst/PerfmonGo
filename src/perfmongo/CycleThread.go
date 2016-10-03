package perfmongo

import (
	"sync"
	"time"
)

type TCycleThread struct {
	Interval    time.Duration
	Function    func()
	ticks       chan bool
	stopChannel chan bool
	active      bool
	waiter      sync.WaitGroup
}

func (this *TCycleThread) tick() {
	for {
		time.Sleep(this.Interval)
		this.ticks <- true
	}
}

func (this *TCycleThread) Start() {
	this.ticks = make(chan bool)
	this.stopChannel = make(chan bool)
	this.active = true
	go this.tick()
	go func() {
		this.waiter.Add(1)
		this.run()
		this.waiter.Done()
	}()
}

func (this *TCycleThread) run() {
	for this.active {
		select {
		case <-this.ticks:
			this.Function()
		case <-this.stopChannel:
			this.active = false
		}
	}
}

// Stop and wait
func (this *TCycleThread) Stop() {
	this.stopChannel <- true
	this.Wait()
}

func (this *TCycleThread) Wait() {
	this.waiter.Wait()
}
