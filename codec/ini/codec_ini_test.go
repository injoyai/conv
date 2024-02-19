package ini

import (
	"fmt"
	"testing"
)

func TestIni_Marshal(t *testing.T) {
	{
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
	{
		i := Ini{}
		bs, err := i.Marshal(map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
				"f": "g",
			},
			"b": map[string]interface{}{
				"b": "c",
				"f": "g",
			},
			"": map[string]interface{}{
				"d": 1,
				"f": 2,
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(string(bs))
	}
}
