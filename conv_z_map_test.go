package conv

import (
	"fmt"
	"testing"
)

func Test_newVal(t *testing.T) {

	s := `{
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

	val := NewValue(s)

	//for i := 0; i < 100; i++ {
	//	go func() {
	t.Log(val.Map())
	t.Log(val.Get("A").Var)
	t.Log(val.Get("A").String())
	fmt.Printf("%#v\n", val.Get("A"))
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
	//	}()
	//}
	//for {
	//}
}
