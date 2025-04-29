package conv

import (
	"testing"
)

func TestGetDefaultString(t *testing.T) {
	t.Log(Default[string]("test"))
	t.Log(Default[string]("test", "a"))
	t.Log(Default[string]("test", "a", "b"))
}

func TestGetDefaultBool(t *testing.T) {
	t.Log(Default[bool](true))
	t.Log(Default[bool](true, false))
	t.Log(Default[bool](true, true))
	t.Log(Default[bool](false, true))
	t.Log(Default[bool](false, false))
	t.Log(Default[bool](false, true, true))
}
