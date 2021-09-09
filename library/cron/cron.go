package cron

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type Job struct {
	Shut chan int `json:"shut"`
}

// implement Run() interface to start rsync job
func (j *Job) Run() {
	log.Printf("your run code")
}

func init() {
	job2 := Job{
		Shut: make(chan int, 1),
	}
	//
	job2.Run()
	//
	go StartJob("* * * * * *", job2)
}

// 定时任务计划
/*
- spec，传入 cron 时间设置
- job，对应执行的任务
*/
func StartJob(spec string, job Job) {
	logger := &CLog{clog: log.New()}
	logger.clog.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	c := cron.New(cron.WithSeconds())

	c.AddJob(spec, &job)

	// 启动执行任务
	c.Start()
	// 退出时关闭计划任务
	defer c.Stop()

	// 如果使用 select{} 那么就一直会循环
	select {
	case <-job.Shut:
		return
	}
}

func StopJob(shut chan int) {
	shut <- 0
}

type CLog struct {
	clog *log.Logger
}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(log.Fields{
		"data": keysAndValues,
	}).Info(msg)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(log.Fields{
		"msg":  msg,
		"data": keysAndValues,
	}).Warn(msg)
}
