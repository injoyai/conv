package conv

import (
	"testing"
)

var s = `{
  "A": [{"a":1,"b":2},{"b":2}],
  "a": "",
  "B": {"b": "bb"},
  "C": [0,1],
  "D": {
    "x": "X",
    "y": "Y"
  },
  "姓名": "钱",
  "l": true,
  "体重": 170.2,
  "年龄": 100,
  "level": {
    "level1": {
      "level2": {
        "level3": "55"
      }
    }
  }
}`

func Test_newVal(t *testing.T) {

	val := NewMap(s)

	//for i := 0; i < 100; i++ {
	//	go func() {
	t.Log(DMap(s).GetString("A"))
	t.Log(GMap(s)["A"])
	t.Log(GMap(s)["B"])
	t.Log(val.Map())
	t.Log(val.Get("A").Var)
	t.Log(val.Get("A").String())
	t.Logf("%#v\n", val.Get("A"))
	t.Log(val.GetString("A[0].a", "ss"))
	t.Log(val.GetString("A[0].b", "ss"))
	t.Log(val.GetString("A[10].a", "ss"))
	t.Log(val.GetString("A[1].a", "ss"))
	t.Log(val.GetString("A[1].b", "ss"))
	t.Log(val.GetString("A[1][1].b", "ss"))
	t.Log(val.GetString("A.a", "ss"))
	t.Log(val.GetString("A", "ss"))
	t.Log(val.Get("A[0]").String("ss"))
	t.Log(val.Get("年龄").Int(20))
	t.Log(val.Get("A").Strings())
	t.Log(val.Get("D.x").String())
	t.Log(val.Get("level.level1.level2.level3").String())
	t.Log(val.Get("level.level1.level2").String())
	t.Log(val.IsDefault("A[1]"))
	t.Log(val.IsDefault())
	//	}()
	//}
	//for {
	//}
}

func Test_Map(t *testing.T) {
	{
		s := `["a","b"]`
		x := NewMap(s)
		t.Log(x.GetString("[1]"))
		if x.GetString("[1]") != "b" {
			t.Error("错误")
		}
	}
	{
		s := `[["c",["d",6]],"b"]`
		x := NewMap(s)
		t.Log(x.GetString("[0][1][1]"))
		if x.GetString("[0][1][1]") != "6" {
			t.Error("错误")
		}
	}
	{
		s := `{"c":["a","b"]}`
		x := NewMap(s)
		t.Log(x.GetString("c[1]"))
		if x.GetString("c[1]") != "b" {
			t.Error("错误")
		}
	}

}

func Test_Map2(t *testing.T) {
	x := NewMap(s)
	a := x.Map()["A"]
	b := a.List()[0]
	t.Log(b.GetString("a"))
	t.Log(a.GetString("[0]"))
	t.Log(a.GetString("[0].a"))

}

// TestNewMap2 [0.67,0.58,0.62,0.59,0.60,0.57]
func TestNewMap2(t *testing.T) {
	for i := 0; i < 100000; i++ {
		NewMap(s)
	}
}
