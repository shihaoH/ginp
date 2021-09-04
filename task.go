package ginp

import (
	"github.com/robfig/cron/v3"
	"sync"
)

// TaskFunc 任务函数定义
type TaskFunc func(params ...interface{})

// TaskExecutor 任务执行者
type TaskExecutor struct {
	f        TaskFunc
	p        []interface{}
	callback func()
}

var (
	once     sync.Once
	taskList chan *TaskExecutor
	onceCron sync.Once
	taskCron *cron.Cron
)

// init 初始化创建协程去跑任务
func init() {
	chlist := getTaskList()
	go func() {
		for exec := range chlist {
			doTask(exec)
		}
	}()
}

// NewTaskExecutor 构造函数
func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

// getTaskList 获取任务列表，单例模式
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

// getCronTask 获取定时任务
func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

// Exec 执行函数
func (t *TaskExecutor) Exec() {
	t.f(t.p...)
}

// doTask 任务协程执行
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

// Task 任务函数
func Task(f TaskFunc, callback func(), p ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, p, callback)
	}()
}
