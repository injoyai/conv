package conv

import (
	"testing"
)

func Test_copySlice(t *testing.T) {
	{
		ls := []int{1, 2, 3}
		ls2 := []string(nil)
		if err := Unmarshal(ls, &ls2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v\n", ls2) //[]string{"1", "2", "3"}
	}
	{
		ls := []int{1, 2, 3}
		ls2 := []string{"0"}
		if err := Unmarshal(ls, &ls2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v\n", ls2) //[]string{"1", "2", "3"}
	}
	{
		ls := New([]int{1, 2, 3})
		ls2 := []string{"0"}
		if err := Unmarshal(ls, &ls2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v\n", ls2)
	}

}

func Test_toInt64(t *testing.T) {
	t.Log(toInt64("-1h1m1s1ms1ns"))
}

func Test_copyMap(t *testing.T) {
	type A struct {
		Name string  `xxx:"name"`
		Age  int     `json:"age"`
		Body bool    `orm:"body"`
		High float64 `json:"high" orm:"high" xxx:"high"`
		Memo string  `json:"age"`
	}
	{ //对象转map
		t.Log("\n===================================对象转Map====================================")
		a := &A{
			Name: "小明",
			Age:  18,
			Body: true,
		}
		{
			m := make(map[string]interface{})
			if err := Unmarshal(a, &m); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%#v\n", m) //map[string]interface {}{"Body":true, "Name":"小明", "age":18}
		}

		{
			m := make(map[string]interface{})
			if err := Unmarshal(a, &m, UnmarshalParam{Tags: []string{"xxx"}}); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%#v\n", m) //map[string]interface {}{"Age":18, "Body":true, "name":"小明"}
		}

		{
			m := make(map[string]interface{})
			if err := Unmarshal(a, &m, UnmarshalParam{Tags: []string{"xxx", "json", "orm"}}); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%#v\n", m) //map[string]interface {}{"Age":18, "Body":true, "name":"小明"}
		}

	}

	{ //map转对象
		t.Log("\n===================================Map转对象====================================")
		m := map[string]interface{}{
			"name":  "小明",
			"age":   18,
			"body":  true,
			"high":  170.6,
			"high1": 170.1,
			"high2": 170.2,
			"high3": 170.3,
			"memo":  "备注信息",
		}
		{
			a := &A{}
			if err := Unmarshal(m, a); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%#v\n", a) //&conv.A{Name:"", Age:18, Body:false, High:170.6, Memo:"18"}
		}
		{
			a := &A{}
			if err := Unmarshal(m, a, UnmarshalParam{Tags: []string{"xxx", "orm", "json"}}); err != nil {
				t.Error(err)
				return
			}
			t.Logf("%#v\n", a) //&conv.A{Name:"小明", Age:18, Body:true, High:170.6, Memo:"18"}
		}

	}

}
