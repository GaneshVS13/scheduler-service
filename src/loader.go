package main

import (
	"context"
	"sync"

	"github.com/scheduler-service/src/config"
	"github.com/scheduler-service/src/scheduler"
)

type StopFunc func()

func loadApplication(ctx context.Context, configFilePath string) (StopFunc, error) {
	wg := &sync.WaitGroup{}

	// load config
	cfg, err := config.Load(configFilePath)
	if err != nil {
		return nil, err
	}

	// load scheduler
	taskScheduler := scheduler.NewTaskScheduler()
	scheduler := scheduler.NewScheduler(cfg.Scheduler, taskScheduler)

	err = scheduler.Load(ctx, wg)
	if err != nil {
		return nil, err
	}

	return func() {
		wg.Wait()
	}, nil
}
