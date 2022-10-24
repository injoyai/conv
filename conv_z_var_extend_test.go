package conv

import (
	"testing"
)

func TestNewExtend(t *testing.T) {
	a := &NewExtendEntity{}
	a.Extend = NewExtend(a)
	t.Log(a.GetString("", "444"))
}

type NewExtendEntity struct {
	*Var
	Extend
}

func (this *NewExtendEntity) GetVar(key string) *Var {
	return New("钱测试")
}
