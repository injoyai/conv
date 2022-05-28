package conv

import "testing"

func TestByte(t *testing.T) {
	t.Log(Byte("00000100"))
	x := falseArray
	x[4] = true
	t.Log(Byte(x))
}
