package cron

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type PeriodicJob struct {
	Ticker time.Ticker
	Job func()
	Cancel context.CancelFunc
}

type PeriodicJobManager struct {
	mu sync.Mutex
	Jobs map[string]*PeriodicJob
}

func NewPeriodicJobManager() *PeriodicJobManager {
	return &PeriodicJobManager{
		Jobs: make(map[string]*PeriodicJob),
	}
}

func (jm *PeriodicJobManager) StartJob(name string, intervalSec float32, job func()) {
	jm.mu.Lock()
	if _, exists := jm.Jobs[name]; exists {
		jm.mu.Unlock()
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	pj := &PeriodicJob{
		Ticker: *time.NewTicker(time.Second * time.Duration(intervalSec)),
		Job: job,
		Cancel: cancel,
	}

	go func() {
		for {
			select {
			case <-pj.Ticker.C:
				pj.Job()
			case <-ctx.Done():
				pj.Ticker.Stop()
				return
			}
		}
	}()

	jm.Jobs[name]=pj
	jm.mu.Unlock()
}

func (jm *PeriodicJobManager) StopJob(name string) error {
	jm.mu.Lock()
	job, exists := jm.Jobs[name]
	if !exists {
		jm.mu.Unlock()
		return fmt.Errorf("job: %s does not exist", name)
	}

	job.Cancel()
	delete(jm.Jobs, name)
	jm.mu.Unlock()

	return nil
}

func (jm *PeriodicJobManager) StopAllJobs() {
	jm.mu.Lock()
	for name, job := range jm.Jobs {
		job.Cancel()
		delete(jm.Jobs, name)
	}

	jm.mu.Unlock()
}
