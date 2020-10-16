package objectpool

import "sync"

type data struct {
	V int64   `json:"v"`
	L []int64 `json:"l"`
}

func newData() *data {
	return &data{
		V: 1,
		L: make([]int64, 0, 10000),
	}
}

// Engine Engine
type Engine struct {
	pool sync.Pool
}

// NewEngine 初始化
func NewEngine() *Engine {
	t := &Engine{}
	t.pool.New = func() interface{} {
		return t.allocate()
	}
	return t
}
func (e *Engine) allocate() *data {
	return newData()
}
