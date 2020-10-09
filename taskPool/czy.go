package taskpool

import "sync"

// PoolCzy PoolCzy
type PoolCzy struct {
	cap            uint64
	runningWorkers uint64
	state          int64
	task           chan *TaskCzy
	sync.Mutex
}

// TaskCzy TaskCzy
type TaskCzy struct {
	Handler func(params ...interface{})
}
