package conv

import (
	"github.com/json-iterator/go"
	"regexp"
	"strings"
)

// NewMap 新建数据
// 递归(惰性,用到才解析)获取所有可以解析数据
func NewMap(i interface{}, model ...string) *Map {
	return newMap(i, model...).decode()
}

func newMap(i interface{}, model ...string) *Map {
	if val, ok := i.(*Map); ok {
		return val
	}
	m := &Map{Var: New(i), model: "json"}
	if len(model) > 0 {
		m.model = model[0]
	}
	return m
}

// Map deep 惰性解析 DMap
// 万次解析0.09s(上版本0.6s)
type Map struct {
	*Var                    //值
	Extend                  //继承
	valMap  map[string]*Map //map
	valList []*Map          //list
	de      bool            //是否解析
	model   string          //解析模式
}

func (this *Map) GoString() string {
	return this.Var.String()
}

func (this *Map) GetVar(key string) *Var {
	return this.Get(key).Var
}

func (this *Map) Map() map[string]*Map {
	for _, v := range this.valMap {
		v.decode()
	}
	return this.valMap
}

func (this *Map) List() []*Map {
	for _, v := range this.valList {
		v.decode()
	}
	return this.valList
}

func (this *Map) Get(key string) *Map {
	data := this
	for _, v := range strings.Split(key, ".") {
		list := regexp.MustCompile(`\[[0-9]+\]`).FindAllString(v, -1)
		lengthAll := 0
		for i, k := range list {
			length := len(k)
			lengthAll += length
			list[i] = k[1 : length-1]
		}
		v = v[:len(v)-lengthAll]
		if len(v) > 0 {
			data = data.getMap(v)
		}
		for _, k := range list {
			data = data.getList(Int(k))
		}
	}
	return data
}

func (this *Map) getMap(key string) *Map {
	data, ok := this.valMap[key]
	if ok {
		return data.decode()
	}
	return NewMap(nil, this.model)
}

func (this *Map) getList(idx int) *Map {
	if idx < len(this.valList) {
		return this.valList[idx].decode()
	}
	return NewMap(nil, this.model)
}

func (this *Map) getParse() func(data []byte, v interface{}) error {
	parse := jsoniter.Unmarshal
	switch this.model {
	case "json":
	case "xml":
	case "toml":
	case "yaml":
	case "ini":
	case "url":
	case "html":
	default:
	}
	return parse
}

func (this *Map) decode() *Map {
	if !this.de {
		parse := this.getParse()
		this.valMap = make(map[string]*Map)
		if !this.Var.IsNil() {
			m := make(map[string]interface{})
			bs := []byte(this.Var.String())
			if err := parse(bs, &m); err == nil {
				for i, v := range m {
					this.valMap[i] = newMap(v, this.model)
				}
			} else {
				var list []interface{}
				if err := parse(bs, &list); err == nil {
					for _, v := range list {
						this.valList = append(this.valList, newMap(v, this.model))
					}
				}
			}
		}
		this.Extend = NewExtend(this)
		this.de = true
	}
	return this
}
