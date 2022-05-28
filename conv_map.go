package conv

import (
	"encoding/json"
	"regexp"
	"strings"
)

type Map struct {
	*Var
	*VarExtend
	valMap   map[string]Map
	valArray []Map
}

// NewValue 新建数据
// 递归获取所有可以解析数据
// @val,任意参数,为nil时也需要声明map空间(否则继承var的方法会panic)
func NewValue(val interface{}) Map {
	data := Map{
		Var:    New(val),
		valMap: make(map[string]Map),
	}
	data.VarExtend = NewVarExtend(data)
	if val != nil {
		m := make(map[string]interface{})
		bs := []byte(String(val))
		if err := json.Unmarshal(bs, &m); err == nil {
			for i, v := range m {
				data.valMap[i] = NewValue(v)
			}
		} else {
			var list []interface{}
			if err := json.Unmarshal(bs, &list); err == nil {
				for _, v := range list {
					data.valArray = append(data.valArray, NewValue(v))
				}
			}
		}
	}
	return data
}

// GoString 格式化"%#v",优先于String()string
func (this Map) GoString() string {
	return this.Var.String()
}

// Get 获取数据,最底层 例如 "a[1].b.c[0]"
func (this Map) Get(key string) Map {
	data := this
	for _, v := range strings.Split(key, ".") {
		var last string
		var ok bool
		if list := regexp.MustCompile(`\[[0-9]+\]`).FindAllString(v, -1); len(list) > 0 {
			last = list[len(list)-1]
			v = v[:len(v)-len(last)]
		}
		if data, ok = data.valMap[v]; !ok {
			data = NewValue(nil)
		} else if length := len(last); length >= 3 {
			if index := Int(last[1 : length-1]); len(data.valArray) > index {
				data = data.valArray[index]
			} else {
				data = NewValue(nil)
			}
		}
	}
	return data
}

// GetVar 实现接口继承方法
func (this Map) GetVar(key string) *Var {
	return this.Get(key).Var
}

func (this Map) Map() map[string]Map {
	return this.valMap
}

func (this Map) Array() []Map {
	return this.valArray
}
