package conv

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log(New(nil))
	t.Log(New(Nil()))
	t.Log(String([]string{"a", "b"}))
	t.Log(New("0xa0").Int())
	t.Log(String(New(100)))
	t.Log(Interfaces([]*Var{
		New(1), New("s"),
	}))
	v := New([]byte{1, 2})
	v2 := v.Copy()
	v.Set("666")
	t.Log(v)
	t.Log(v2)
	t.Log(v.String())
	t.Log(v2.String())
}
