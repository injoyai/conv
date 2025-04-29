package conv

import (
	"testing"
)

func TestNewExtend(t *testing.T) {
	a := &NewExtendEntity{}
	a.ExtendString = NewExtendDefault(a)
	t.Log(a.GetString("", "444"))
}

type NewExtendEntity struct {
	*Var
	ExtendString
}

func (this *NewExtendEntity) GetVar(key string) *Var {
	return New("钱测试")
}
