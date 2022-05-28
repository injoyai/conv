package conv

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log(String([]string{"a", "b"}))
	t.Log(New("0xa0").Int())
}
