package core

import (
	"context"
	"errors"
	"math"
	"sync"
)

const BUF_STAT_APPLY_DISABLE = 0x1
const BUF_STAT_FREE_DISABLE = 0x2

type Buffer struct {
	id int

	// config
	bufSize int64

	charSet *CharSet

	// runtime
	// Do not use channel because it is inconvenient to random-read
	buf  chan string
	mBuf sync.Map

	stat uint8

	ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

var curBufferId = 0

func NewBuffer() (b *Buffer, err error) {
	ctx, ctxCancelFunc := context.WithCancel(context.Background())
	b = &Buffer{
		id:            curBufferId,
		bufSize:       BufferConfig.BufSize,
		charSet:       NewCharSet(BufferConfig.CharSet, BufferConfig.MinLength, BufferConfig.IncrLength),
		buf:           make(chan string, BufferConfig.BufSize),
		ctx:           ctx,
		ctxCancelFunc: ctxCancelFunc,
	}
	curBufferId++
	buf := make([]string, b.bufSize)

	for i := int64(0); i < b.bufSize; i++ {
		bs, err := b.charSet.NextId()
		if err != nil {
			return nil, err
		}
		s := string(bs)
		buf[i] = s
		b.mBuf.Store(s, false)
	}

	Shuffle(buf, 0, b.bufSize-1)

	for _, t := range buf {
		b.buf <- t
	}

	return b, nil
}

// err != nil when the new-token quantity less than cnt
// err != nil when added token will not be remove
func (b *Buffer) Fill(cnt uint8) (err error) {
	buf := make([]string, cnt)
	var i uint8 = 0
	var bs []byte
	for i < cnt {
		bs, err = b.charSet.NextId()
		if err != nil {
			return
		}
		buf[i] = string(bs)
		i++
		if b.bufSize < math.MaxInt64 {
			b.bufSize++
		}
	}
	return
}

func (b *Buffer) ApplyCh() <-chan string {

	return b.buf
}

func (b *Buffer) Apply( /*ctx context.Context*/) (token string /*, err error*/) {
	//if ctx == nil {
	token = <-b.buf
	//	b.mBuf.Store(token, false)
	//} else {
	//	select {
	//	case <-ctx.Done():
	//		err = errors.New(`there is no token to apply for`)
	//	case token = <-b.buf:
	//		b.mBuf.Store(token, false)
	//	}
	//}

	return
}

func (b *Buffer) Free(token string) (err error) {
	if _, ok := b.mBuf.Load(token); !ok {
		return errors.New(`not found id: ` + token)
	}
	//b.mBuf.Store(token, true)
	b.buf <- token

	return
}

func (b *Buffer) CheckExist(token string) (ok bool) {
	_, ok = b.mBuf.Load(token)
	return
}

func (b *Buffer) isCanApply() bool {
	return b.stat&BUF_STAT_APPLY_DISABLE > 0 && len(b.buf) > 0
}
