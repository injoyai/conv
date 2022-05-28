package conv

import (
	"encoding/hex"
	"testing"
)

func TestBytes(t *testing.T) {
	t.Log(hex.EncodeToString(Bytes(100)))
}
