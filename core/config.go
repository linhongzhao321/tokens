package core

import "time"

var BufferConfig = struct {
	BufSize      int64
	CharSet      string
	MinLength    uint8
	IncrLength   uint8
	BufCnt       int
	ApplyTimeout time.Duration
}{
	BufSize:      1000000,
	CharSet:      `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`,
	MinLength:    4,
	IncrLength:   8,
	BufCnt:       2,
	ApplyTimeout: time.Second,
}
