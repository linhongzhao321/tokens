package test

import (
	"github.com/linhongzhao321/tokens/core"
	"testing"
	"time"
)

func BenchmarkChannal(b *testing.B) {
	ch := make(chan struct{}, 50000000)
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = <-ch
	}
}

func TestShuffle(t *testing.T) {
	var begin, end int64
	begin = time.Now().UnixNano()
	s := make([]string, 10000000)
	end = time.Now().UnixNano()
	t.Log(`init spend: `, end-begin, `ns `)
	begin = time.Now().UnixNano()
	core.Shuffle(s, 0, 9999999)
	end = time.Now().UnixNano()
	t.Log(`shuffle spend: `, end-begin, `ns`)
}

func BenchmarkPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log("hello")
	}
}