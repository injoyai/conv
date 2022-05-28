package conv

import (
	"testing"
)

func TestNewExtend(t *testing.T) {
	a := &NewExtendEntity{}
	a.VarExtend = NewVarExtend(a)
	t.Log(a.GetString("", "444"))
}

type NewExtendEntity struct {
	*Var
	*VarExtend
}

func (this *NewExtendEntity) GetVar(key string) *Var {
	return New("钱测试")
}
