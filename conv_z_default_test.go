package conv

import (
	"testing"
)

func TestGetDefaultString(t *testing.T) {
	t.Log(GetDefaultString("test"))
	t.Log(GetDefaultString("test", "a"))
	t.Log(GetDefaultString("test", "a", "b"))
}

func TestGetDefaultBool(t *testing.T) {
	t.Log(GetDefaultBool(true))
	t.Log(GetDefaultBool(true, false))
	t.Log(GetDefaultBool(true, true))
	t.Log(GetDefaultBool(false, true))
	t.Log(GetDefaultBool(false, false))
	t.Log(GetDefaultBool(false, true, true))
}
