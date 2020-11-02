package test

import (
	"github.com/linhongzhao321/tokens/core"
	"testing"
)

//func TestCharSetZeroMinBytes(t *testing.T) {
//	tokens.NewCharSet()
//}

const noErrMsg = `No error should be reported under the current parameter. the actual error:`

func TestCharSetAllEmpty(t *testing.T) {
	cs := core.NewCharSet(``, 0, 0)
	id, err := cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if id == nil {
		t.Error(`returned id should not be nil, the actual value is nil`)
		return
	}
	if len(id) > 0 {
		t.Error(`returned length of id([]byte) should be 0, the actual value is`, len(id))
	}
}

func TestCharSetStringEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(`here  should caught error[bs should not empty!]`)
		} else if r.(string) != `bs should not empty!` {
			t.Error(`the acture panic: '`, r, `'`)
		}
	}()
	_ = core.NewCharSet(``, 1, 3)
}

func TestCharSetZeroBytes(t *testing.T) {
	cs := core.NewCharSet(`1234567890123456`, 0, 0)
	id, err := cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if id == nil {
		t.Error(`returned id should not be nil, the actual value is nil`)
		return
	}
	if len(id) > 0 {
		t.Error(`returned length of id([]byte) should be 0, the actual value is`, len(id))
	}
}

func TestCharSetZeroMinBytes(t *testing.T) {
	cs := core.NewCharSet(`1234567890123456`, 0, 3)
	id, err := cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 0 {
		t.Error(`The first ID length should be 0. the actual length is`, len(id))
	}
	id, err = cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 1 {
		t.Error(`The first ID length should be 0. the actual length is`, len(id))
	}
}

func TestCharSetStringLengthIs1(t *testing.T) {
	cs := core.NewCharSet(`a`, 0, 3)
	id, err := cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 0 {
		t.Error(`The ID length should be 0. the actual length is`, len(id))
	}
	//t.Log(`id: '`+string(id) +`'`)
	id, err = cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 1 {
		t.Error(`The ID length should be 0. the actual length is`, len(id))
	}
	if string(id) != `a` {
		t.Error(`The ID length should be 'a'. the actual:`, len(id))
	}
	//t.Log(`id: '`+string(id) +`'`)
	id, err = cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 2 {
		t.Error(`The ID length should be 0. the actual length is`, len(id))
	}
	if string(id) != `aa` {
		t.Error(`The ID length should be 'aa'. the actual:`, len(id))
	}
	//t.Log(`id: '`+string(id) +`'`)
	id, err = cs.NextId()
	if err != nil {
		t.Error(noErrMsg, err.Error())
		return
	}
	if len(id) != 3 {
		t.Error(`The ID length should be 0. the actual length is`, len(id))
	}
	if string(id) != `aaa` {
		t.Error(`The ID length should be 'aaa'. the actual:`, len(id))
	}
	//t.Log(`id: '`+string(id) +`'`)
}

func BenchmarkCharSetBase(b *testing.B) {
	cs := core.NewCharSet(`1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`, 4, 12)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := cs.NextId()
		if err != nil {
			b.Error(`err:`, err, `have run: `, i)
			return
		}
	}
}

func TestCharSetBase(t *testing.T) {
	cs := core.NewCharSet(`1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`, 4, 0)

	for i := 0; i < 14776336; i++ {
		_, err := cs.NextId()
		if err != nil {
			t.Error(noErrMsg, err)
			return
		}
	}

	_, err := cs.NextId()
	if err == nil {
		t.Error(`here should have error that is 'generated numeric has max''`)
	}
}
