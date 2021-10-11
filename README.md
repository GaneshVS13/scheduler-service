# Scheduler Service
This is a configurable cron sceduler service which performs given task in scheduled time (in minutes).

# Packages
Used `github.com/jasonlvhit/gocron` package to configure the cron schedule.
Used `github.com/go-redis/redis` for locking mechanism over cron schedule.