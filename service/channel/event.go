package channel

import (
	"fmt"
	"time"
)

type event struct {
	running chan struct{}
}

func (ev *event) Start() error {
	if ev.IsRunning() {
		return fmt.Errorf("event is running")
	}
	started := make(chan error)
	go func() {
		ev.running = make(chan struct{})
		defer close(ev.running)
		started <- fmt.Errorf("recreate event client failed")
		started <- fmt.Errorf("register filtered block event failed")
		close(started)

		time.Sleep(100 * time.Second)
		fmt.Println("========================")
	}()
	return <-started
}

func (ev *event) IsRunning() bool {
	if ev.running != nil {
		select {
		case <-ev.running:
		default:
			return true
		}
	}
	return false
}
