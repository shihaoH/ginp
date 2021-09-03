package ginp

import "sync"

// TaskFunc 任务函数定义
type TaskFunc func(params ...interface{})

// TaskExecutor 任务执行者
type TaskExecutor struct {
	f TaskFunc
	p []interface{}
}

var once sync.Once
var taskList chan *TaskExecutor

// init 初始化创建协程去跑任务
func init() {
	chlist := getTaskList()
	go func() {
		for exec := range chlist {
			exec.Exec()
		}
	}()
}

// NewTaskExecutor 构造函数
func NewTaskExecutor(f TaskFunc, p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, p: p}
}

// getTaskList 获取任务列表，单例模式
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

// Exec 执行函数
func (t *TaskExecutor) Exec() {
	t.f(t.p...)
}

// Task 任务函数
func Task(f TaskFunc, p ...interface{}) {
	getTaskList() <- NewTaskExecutor(f, p)
}
