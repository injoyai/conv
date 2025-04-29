package conv

import (
	"testing"
	"time"
)

func TestIsDefault(t *testing.T) {
	t.Log(IsDefault(0))
	t.Log(IsDefault(""))
	t.Log(IsDefault(struct{ Name string }{}))
	t.Log(IsDefault(make(map[string]string)))
	t.Log(IsDefault(map[string]any{"name": "test"}))
	t.Log(IsDefault([]string{}))
	t.Log(IsDefault([]string{"a"}))
	var a []string
	t.Log(IsDefault(a))
	var m map[string]any
	t.Log(IsDefault(m))
	t.Log(IsDefault(&map[string]any{"name": "test"}))
	t.Log(IsDefault(&map[string]any{}))
	t.Log(IsDefault(map[string]any{}))

	t.Log(IsDefault(struct{ Name string }{}))
	t.Log(IsDefault(struct {
		Name string
		Age  int
	}{}))
	t.Log(IsDefault(struct {
		Name string
		Age  int
	}{
		Name: "",
	}))
	t.Log(IsDefault(struct {
		Name string
		Age  int
	}{
		Name: "1",
	}))

}

func TestIsList(t *testing.T) {
	var x **[]string
	y := []string{}
	z := &y
	x = &z
	_ = x
	t.Log(IsArray(x))
}

func TestGetNature(t *testing.T) {
	var x **[]string
	y := []string{}
	z := &y
	x = &z
	_ = x
	t.Log(GetNature(nil))
}

func TestIsTime(t *testing.T) {
	var x *time.Time
	y := time.Now()
	x = &y
	//x = nil
	t.Log(IsTime(x))
}

func TestIsNumber(t *testing.T) {
	var x *int
	y := 1.6
	z := int(y)
	x = &z
	t.Log(IsNumber(x))
}
