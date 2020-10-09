package objectpool

import (
	"testing"
)

func Benchmark_newEngine1(b *testing.B) {
	n := NewEngine()
	for index := 0; index < b.N; index++ {

		t := n.pool.Get().(*data)
		t1 := n.pool.Get().(*data)
		t2 := n.pool.Get().(*data)
		t.V = 2
		n.pool.Put(t)
		n.pool.Put(t1)
		n.pool.Put(t2)
	}
}
func Benchmark_newEngineWithoutPool(b *testing.B) {
	for index := 0; index < b.N; index++ {
		t := newData()
		t.V = 2
		t1 := newData()
		t1.V = 2
		t2 := newData()
		t2.V = 2
	}
}
