// 不同 short id 生成方案对比
package test

import (
	"context"
	"fmt"
	"github.com/linhongzhao321/tokens/core"
	"testing"
	"time"
)

var buf *core.Buffer

var initBuf = true

func init() {
	core.BufferConfig.BufSize = 10000000
	core.BufferConfig.CharSet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	if initBuf {
		beginTime := time.Now().UnixNano()
		var err error
		buf, err = core.NewBuffer()
		if err != nil {
			panic(err.Error())
			return
		}
		if buf == nil {
			panic(`token buffer is empty`)
			return
		}
		useTime := time.Now().UnixNano() - beginTime
		fmt.Println(
			`[buffer size]`, core.BufferConfig.BufSize,
			` [use time(ns)]`, useTime,
			` [use time(ms)]`, useTime/int64(time.Millisecond),
			` [use time(s)]`, useTime/int64(time.Second),
		)

	}
}

// [6, 8]s
func TestCreate(t *testing.T) {
	beginTime := time.Now().UnixNano()
	b, err := core.NewBuffer()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if b == nil {
		t.Error(`token buffer is empty`)
		return
	}
	useTime := time.Now().UnixNano() - beginTime
	t.Log(
		`[buffer size]`, core.BufferConfig.BufSize,
		` [use time(ns)]`, useTime,
		` [use time(ms)]`, useTime/int64(time.Millisecond),
		` [use time(s)]`, useTime/int64(time.Second),
	)
	return
}

// [6, 8]s
func TestApply(t *testing.T) {
	beginTime := time.Now().UnixNano()
	_ = buf.Apply()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	useTime := time.Now().UnixNano() - beginTime
	t.Log(useTime)
	return
}

func TestReadCh(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Microsecond)
	select {
	case <-ctx.Done():
		t.Error("timeout")
	case token := <-buf.ApplyCh():
		t.Log(token)
	}
	return
}

// 100, 200 ns
func BenchmarkReadCh(b *testing.B) {
	fmt.Println("running...")
	//bg := context.Background()
	//var token string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		token := <-buf.ApplyCh()
		b.StopTimer()
		err := buf.Free(token)
		if err != nil {
			b.Error(err)
			return
		}
	}
	return
}

// [6, 8]s
func TestFree(t *testing.T) {
	token := <-buf.ApplyCh()
	err := buf.Free(token)
	if err != nil {
		t.Error(err)
	}
	return
}

// 运行前，请先依据当前机器性能，在init设置足够的BufferConfig.BufSize
func BenchmarkApply(b *testing.B) {
	b.Log("benchmark running...")
	//var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = buf.Apply()
		//if err != nil {
		//	b.Error(err.Error(), `has run:`, i+1)
		//	return
		//}
	}
}

// 运行前，请先依据当前机器性能，在init设置足够的BufferConfig.BufSize
// 600, 700 ns
func BenchmarkFree(b *testing.B) {
	fmt.Println("running...")
	//bg := context.Background()
	//var token string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token := <-buf.ApplyCh()
		b.StartTimer()
		err := buf.Free(token)
		b.StopTimer()
		if err != nil {
			b.Error(err)
			return
		}
	}
	return
}
