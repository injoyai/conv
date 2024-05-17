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
		t.Logf("%#v\n", ls2)
	}
	{
		ls := []int{1, 2, 3}
		ls2 := []string{"0"}
		if err := Unmarshal(ls, &ls2); err != nil {
			t.Error(err)
			return
		}
		t.Logf("%#v\n", ls2)
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
