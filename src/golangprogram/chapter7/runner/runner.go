package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("timeout error")
	ErrInterrupt = errors.New("interrupt error")
)

type Runner struct {
	interruptChan chan os.Signal   // 中断信号通道
	completeChan  chan error       // 完成通道
	timeOutChan   <-chan time.Time // 超时通道
	tasks         []func(int)      // 执行任务列表
}

func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interruptChan: make(chan os.Signal, 1),
		completeChan:  make(chan error),
		timeOutChan:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interruptChan, os.Interrupt)

	go func() {
		r.completeChan <- r.run()
	}()

	select {
	case err := <-r.completeChan:
		return err
	case <-r.timeOutChan:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.GotInterrupt() {
			return ErrInterrupt
		} else {
			task(id)
		}
	}
	return nil
}

func (r *Runner) GotInterrupt() bool {
	select {
	case <-r.interruptChan:
		signal.Stop(r.interruptChan)
		return true
	default:
		return false
	}
}
