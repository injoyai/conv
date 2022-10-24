package conv

import (
	"encoding/binary"
	"testing"
)

func Test_padding(t *testing.T) {
	bytes := []byte{0x00, 0x01}
	number := binary.BigEndian.Uint64(bytes)
	t.Log(number)
}
