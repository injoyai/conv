package ini

import (
	"fmt"
	"strings"
	"testing"
)

func TestIni_Marshal(t *testing.T) {
	i := Ini{}
	bs, err := i.Marshal(map[string]map[string]string{
		"a": {
			"b": "c",
			"f": "g",
		},
		"b": {
			"b": "c",
			"f": "g",
		},
		"": {
			"d": "e",
			"f": "g",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bs))
}

func TestIni_Unmarshal_InvalidTargetType(t *testing.T) {
	i := Ini{}
	target := map[string]string{}

	err := i.Unmarshal([]byte("a=b\n"), &target)
	if err == nil {
		t.Fatal("expected error for invalid target type")
	}
	if !strings.Contains(err.Error(), "*map[string]any") {
		t.Fatalf("expected error to mention expected type, got %v", err)
	}
	if !strings.Contains(err.Error(), "*map[string]string") {
		t.Fatalf("expected error to mention actual type, got %v", err)
	}
}
