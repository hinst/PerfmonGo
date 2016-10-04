package perfmongo

import (
	"os"
	"os/signal"
)

func InstallShutdownReceiver(receiver func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	go func() {
		for currentSignal := range c {
			var _ = currentSignal
			receiver()
		}
	}()
}
