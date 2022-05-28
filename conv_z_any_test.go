package conv

import (
	"encoding/hex"
	"testing"
)

func TestBytes(t *testing.T) {
	t.Log(hex.EncodeToString(Bytes(100)))
	t.Log(Byte("00000100"))
	x := falseArray
	x[4] = true
	t.Log(Byte(x))
}
