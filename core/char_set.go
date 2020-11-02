// 本模块按照指定字符集
// 依次生成该字符集中所有字符存在的组合情况
// 通过 NewCharSet() 初始化
// 通过 CharSet.NextId() 每次返回一个
package core

import (
	"errors"
)

type CharSet struct {
	set    string
	setCnt int64

	maxBytes uint16
	minBytes uint16
	// 每次生成由curBytes 的 curNum + 1 开始生成
	// 当 curNum > maxNum 时, curBytes + 1, curNum = 0
	maxNum   []int64
	curBytes uint16
	curNum   int64
}

func NewCharSet(bs string, minBytes uint8, incrBytes uint8) *CharSet {
	if len(bs) == 0 {
		panic("bs should not empty!")
	}

	incrBytesU16 := uint16(incrBytes)
	minBytesU16 := uint16(minBytes)
	bsLen := int64(len(bs))
	cs := &CharSet{
		set:      bs,
		setCnt:   bsLen,
		maxBytes: minBytesU16 + incrBytesU16,
		minBytes: incrBytesU16,
		curBytes: minBytesU16,
		curNum:   -1,
	}
	cs.maxNum = make([]int64, cs.maxBytes+2)
	var t int64
	maxBytes64 := int64(cs.maxBytes)
	for t = 0; t <= maxBytes64; t++ {
		cs.maxNum[t] = IntPow(bsLen, t)
	}
	cs.maxNum[t] = IntPow(bsLen, t)
	return cs
}

func (cSet *CharSet) NextId() ([]byte, error) {
	// generate cur-num
	cSet.curNum += 1
	if cSet.curNum >= cSet.maxNum[cSet.curBytes] {
		if cSet.curBytes >= cSet.maxBytes {
			return nil, errors.New("generated numeric has max")
		}
		cSet.curNum = 0
		cSet.curBytes += 1
	}

	// to []byte
	buf := make([]byte, cSet.curBytes)
	curNum := cSet.curNum
	for i := cSet.curBytes; i > 0; i-- {
		t := curNum % cSet.setCnt
		curNum /= cSet.setCnt
		buf[i-1] = cSet.set[t]
	}

	return buf, nil
}
