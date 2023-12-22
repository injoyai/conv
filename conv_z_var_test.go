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
}
