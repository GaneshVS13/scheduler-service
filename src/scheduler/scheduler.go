package scheduler

import (
	"context"
	"log"
	"sync"

	"github.com/jasonlvhit/gocron"

	"github.com/scheduler-service/src/config"
)

// Job is an entity to hold job information
type Job struct {
	schedule uint64
	task     func(context.Context)
}

// Scheduler is an entity of scheduler
type Scheduler struct {
	cfg           config.Scheduler
	taskScheduler *TaskScheduler
	cronScheduler *gocron.Scheduler
}

// NewScheduler returns instance of scheduler
func NewScheduler(
	c config.Scheduler,
	ts *TaskScheduler,
) *Scheduler {
	// locker implementation
	// locker := &locker{
	// 	redisClient,
	// }
	// gocron.SetLocker(locker)

	return &Scheduler{
		cfg:           c,
		taskScheduler: ts,
		cronScheduler: gocron.NewScheduler(),
	}
}

func (sc *Scheduler) Load(ctx context.Context, wg *sync.WaitGroup) error {
	jobs := make([]Job, 0)

	jobs = append(jobs,
		Job{
			schedule: sc.cfg.TaskSchedule,
			task:     sc.taskScheduler.Run,
		},
	)

	return sc.initScheduler(ctx, jobs, wg)
}

func (sc *Scheduler) initScheduler(ctx context.Context, jobs []Job, wg *sync.WaitGroup) error {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				sc.clearSchedules()
				return
			}
		}
	}()

	err := sc.startScheduler(ctx, jobs)
	if err != nil {
		return err
	}

	return nil
}

func (sc *Scheduler) startScheduler(ctx context.Context, jobs []Job) error {
	log.Println("Starting scheduler...")

	for _, job := range jobs {
		// uncomment below line while using locker
		//err := sc.cronScheduler.Every(job.schedule).Seconds().Lock().Do(job.task, ctx)
		err := sc.cronScheduler.Every(job.schedule).Minutes().Do(job.task, ctx)
		if err != nil {
			return err
		}
	}

	sc.cronScheduler.Start()

	return nil
}

func (sc *Scheduler) clearSchedules() {
	log.Println("Stopping scheduler...")
	sc.cronScheduler.Clear()
}
