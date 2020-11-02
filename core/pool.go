package core

import (
	"context"
)

type BufferPool struct {
	ctx    context.Context
	bufs   map[int]*Buffer
	bufCnt int
	//locks   map[int]*sync.Mutex
	usedCnt int
	notUsed chan *Buffer
}

func NewBufferPool() (*BufferPool, error) {
	bp := &BufferPool{
		bufCnt:  BufferConfig.BufCnt,
		notUsed: make(chan *Buffer, BufferConfig.BufCnt),
		bufs:    make(map[int]*Buffer),
	}

	for i := 0; i < BufferConfig.BufCnt; i++ {
		b, err := NewBuffer()
		if err != nil {
			return nil, err
		}
		bp.bufs[i] = b
		bp.notUsed <- b
	}

	return bp, nil
}

func (bp *BufferPool) buffer() (b *Buffer) {
	b = <-bp.notUsed
	return
}

func (bp *BufferPool) bufferByBufId(bufId int) (b *Buffer) {
	b = <-bp.notUsed
	return
}

func (bp *BufferPool) freeBuffer(b *Buffer) {
	bp.notUsed <- b
}

func (bp *BufferPool) Apply() (bufId int, token string, err error) {
	b := bp.buffer()
	token = b.Apply()
	bp.freeBuffer(b)

	return b.id, token, err
}

func (bp *BufferPool) Free(bufId int, token string) (err error) {
	b := bp.bufferByBufId(bufId)
	err = b.Free(token)
	bp.freeBuffer(b)

	return err
}
