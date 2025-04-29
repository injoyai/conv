package conv

import (
	"github.com/injoyai/conv/codec"
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
	{
		s := `{"c":["a","b","c","d"]}`
		x := NewMap(s)
		x.Set("c[1]", "bb")
		t.Log(x.String())
		x.Set("c[100]", "c100")
		t.Log(x.String())
		x.Set("c[3]", `{"a":1,"b":2}`)
		t.Log(x.String())
		t.Log(x.GetString("c[3].a"))
		t.Log(x.Get("c").String())
		x.Set("b", `{"a":1,"b":2}`)
		t.Log(x.String())
		x.Set("c[2]", map[string]any{
			"aa": 1,
			"bb": true,
			"cc": 1.02,
			"dd": []any{12, 10.3, false, true, "999"},
			"ee": map[string]any{
				"aaa": map[any]any{
					100: 101, "102": 103, 104.1: 105, true: false,
				},
			},
		})
		t.Log(x.String())
		t.Log(x.GetString("c[2].dd[1]"))
		t.Log(x.GetString("c[2].ee.aaa"))
		t.Log(x.GetString("c[2].ee.aaa.100"))
		t.Log(x.GetString("c[2].ee.aaa.104.1")) //小数点不支持,和分隔符冲突了
		t.Log(x.GetString("c[2].ee.aaa.true"))
	}

	{
		s := `{"c":["a","b","c","d"]}`
		x := NewMap(s)
		x.Append("c", "e", "ee", "eee")
		t.Log(x.String())
		t.Log(x.GetString("c"))
	}

}

func Test_Map3(t *testing.T) {
	{
		s := `{"c":["a","b","c","d"]}`
		x := NewMap(s)
		x.Append("c", "e", "ee", "eee")
		t.Log(x.String())
		t.Log(x.GetString("c"))
		t.Log(x.GetString("c[6]"))
		x.Append("d", 0, 1, 2, 3)
		t.Log(x.String())
		x.Append("e")
		x.Append("f", 0)
		t.Log(x.String())
		x.Append("e", 1)
		t.Log(x.String())
		t.Log(x.GetString("e[0]"))
		x.Set("f", map[string]any{"aa": 1, "bb": 2})
		t.Log(x.String())
		x.Append("f", 1, 2, 3)
		x.Set("g", nil)
		t.Log(x.String())
		x.Del("c[0]")
		x.Del("c[-1]")
		x.Del("d[0]")
		x.Del("d[-1]")
		x.Del("e[0]")
		x.Del("d[10]")
		x.Del("f.aa")
		x.Del("g")
		t.Log(x.String())
		x.Del("c")
		t.Log(x.String())
		x.Set("x", 100)
		t.Log(x.String())
		x.Set("y.y[0].y", "sss")
		t.Log(x.String())
		x.Set("", 199)
		t.Log(x.String())
	}
}

func Test_Map2(t *testing.T) {
	x := NewMap(s)
	list := x.Get("A")
	object := list.Get("[0]")
	t.Log(object.GetString("a"))   //"1"
	t.Log(object.GetString("[0]")) //""
	t.Log(list.GetString("[0]"))   //{"a":1,"b":2}
	t.Log(list.GetString("[0].a")) //"1"

}

func Test_Map4(t *testing.T) {
	yaml := `name: yaml
age: 18`
	x := NewMap(yaml, codec.Yaml)
	t.Log(x.GetInt("age"))
	t.Log(x.String())
	x.Set("high", 180)
	t.Log(x.String())
}

// TestNewMap2 [0.67,0.58,0.62,0.59,0.60,0.57]
func TestNewMap2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		NewMap(s)
	}
}
