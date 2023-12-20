package conv

import (
	"github.com/injoyai/conv/codec"
	json "github.com/json-iterator/go"
	"regexp"
	"strings"
)

// NewMap 新建数据
// 递归(惰性,用到才解析)获取所有可以解析数据
func NewMap(i interface{}, codec ...codec.Interface) *Map {
	return newMap(i, codec...).decode()
}

func newMap(i interface{}, codec ...codec.Interface) *Map {
	if val, ok := i.(*Map); ok {
		return val
	}
	m := &Map{Var: New(i)}
	if len(codec) > 0 {
		m.codec = codec[0]
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
	codec   codec.Interface //编解码
}

func (this *Map) IsDefault(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsDefault()
	}
	return this.Get(key[0]).IsDefault(key[1:]...)
}

func (this *Map) IsNumber(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsNumber()
	}
	return this.Get(key[0]).IsNumber(key[1:]...)
}

func (this *Map) IsString(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsString()
	}
	return this.Get(key[0]).IsString(key[1:]...)
}

func (this *Map) IsInt(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsInt()
	}
	return this.Get(key[0]).IsInt(key[1:]...)
}

func (this *Map) IsFloat(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsFloat()
	}
	return this.Get(key[0]).IsFloat(key[1:]...)
}

func (this *Map) IsBool(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsBool()
	}
	return this.Get(key[0]).IsBool(key[1:]...)
}

func (this *Map) IsTime(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsTime()
	}
	return this.Get(key[0]).IsTime(key[1:]...)
}

func (this *Map) IsArray(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsArray()
	}
	return this.Get(key[0]).IsArray(key[1:]...)
}

func (this *Map) IsPointer(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsPointer()
	}
	return this.Get(key[0]).IsPointer(key[1:]...)
}

func (this *Map) IsNil(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsNil()
	}
	return this.Get(key[0]).IsNil(key[1:]...)
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
	return NewMap(nil, this.codec)
}

func (this *Map) getList(idx int) *Map {
	if idx < len(this.valList) {
		return this.valList[idx].decode()
	}
	return NewMap(nil, this.codec)
}

func (this *Map) getParse() func(data []byte, v interface{}) error {
	if this.codec == nil {
		if codec.Default != nil {
			return codec.Default.Unmarshal
		}
		return json.Unmarshal
	}
	return this.codec.Unmarshal
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
					this.valMap[i] = newMap(v, codec.Json)
				}
			} else {
				var list []interface{}
				if err := parse(bs, &list); err == nil {
					for _, v := range list {
						this.valList = append(this.valList, newMap(v, codec.Json))
					}
				}
			}
		}
		this.Extend = NewExtend(this)
		this.de = true
	}
	return this
}
