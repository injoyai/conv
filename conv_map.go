package conv

import (
	"github.com/injoyai/conv/codec"
	json "github.com/json-iterator/go"
	"regexp"
	"strings"
)

// NewMap 新建数据
// 递归(惰性,用到才解析)获取所有可以解析数据
func NewMap(i any, codec ...codec.Interface) *Map {
	//一般用NewMap之后,肯定会获取值的,固直接decode
	return newMap(i, codec...).decode()
}

func newMap(i any, codec ...codec.Interface) *Map {
	if val, ok := i.(*Map); ok {
		return val
	}
	m := &Map{Var: New(i)}
	if len(codec) > 0 && codec[0] != nil {
		m.codec = codec[0]
	}
	return m
}

// Map deep 惰性解析 DMap
// 万次解析0.09s
type Map struct {
	*Var                    //值
	Extend                  //继承
	isArray bool            //是否是对象,否则是列表,用于处理内容为空时,显示[]还是{}的情况
	object  map[string]*Map //对象实例
	array   []*Map          //数组实例
	codec   codec.Interface //编解码
	decoded bool            //惰性加载,解析一次
	hasSet  bool            //设置了新数据
}

// GetVar 实现Extend的接口,继承Extend的接口
func (this *Map) GetVar(key string) *Var {
	return this.Get(key).Var
}

// Map 获取对象map,不安全,不推荐使用
func (this *Map) Map() map[string]*Map {
	for _, v := range this.object {
		v.decode()
	}
	return this.object
}

// List 获取数组,不安全,不推荐使用
func (this *Map) List() []*Map {
	for _, v := range this.array {
		v.decode()
	}
	return this.array
}

// Get 获取Map,可以根据分隔符来获取更深入的数据
// 分隔符.和小数冲突,现在暂时不支持查询key是小数的对象或者key有.和[]
// 例如使用key(a.b[0]) 可以从 {"a":{"b":[0,1]}} 得到0的Map
func (this *Map) Get(key string) *Map {
	data := this
	for _, v := range this.getKeys(key) {
		switch k := v.(type) {
		case string:
			data = data.getObject(k, false)
		case int:
			data = data.getArray(k)
		}
	}
	return data.refresh()
}

// Set 设置数据,会覆盖原先的数据,
// key可以使用分割符号.和[],分别对应对象和数组,例 a.b[1].c[1][3]
// 分隔符.和小数冲突,现在暂时不支持查询key是小数的对象或者key有.和[]
// 例如使用Set("a.b[0]",100) 可以使 {"a":{"b":[0,1]}} 变成 {"a":{"b":[100,1]}}
// 假如对象key不存在,会自动添加
// 如果数组的下标不存在,则不会自动添加,例如设置了下标100,补充99个nil值不是很合理
// 如果设置了nil,对应json的null
func (this *Map) Set(key string, value any) *Map {
	data := this
	for _, v := range this.getKeys(key) {
		//所有的父级都打上标记,方便后续判断
		data.hasSet = true
		switch k := v.(type) {
		case string:
			data = data.getObject(k, true)
		case int:
			data = data.getArray(k)
		}
	}
	*data = *newMap(value, this.codec) //覆盖原先的Map
	return this
}

// Append 追加元素,只对数组有效
// key可以使用分割符号.和[],分别对应对象和数组,例 a.b[1].c[1][3]
// 分隔符.和小数冲突,现在暂时不支持查询key是小数的对象或者key有.和[]
// 例如 {"a":{"n":1},"b":[0,1]} 使用Append("b",3,4)后,得到 {"a":{"n":1},"b":[0,1,3,4]}
// 但是使用Append("a",3,4)后是无变化的,因为a是对象,可以使用Set("a",[]int{0,1,3,4})得到一样的结果
// 如果设置了nil,对应json的null
func (this *Map) Append(key string, value ...any) *Map {
	data := this
	for _, v := range this.getKeys(key) {
		//所有的父级都打上标记,方便后续判断
		data.hasSet = true
		switch k := v.(type) {
		case string:
			data = data.getObject(k, true)
		case int:
			data = data.getArray(k)
		}
	}
	if data.isArray || len(data.object) == 0 {
		//假如不是对象的情况下生效,或者对象是空的情况
		data.isArray = true
		for _, v := range value {
			data.array = append(data.array, newMap(v, this.codec))
		}
		data.hasSet = true
		data.Var.Set(data.encode())
	}
	return this
}

func (this *Map) Del(key string) {
	data := this
	keys := this.getKeys(key)
	for i, v := range this.getKeys(key) {
		//所有的父级都打上标记,方便后续判断
		data.hasSet = true
		//当key是最后一个时候,进行删除对象键或者删除数组元素操作
		if i == len(keys)-1 {
			switch k := v.(type) {
			case string:
				delete(data.object, k)
			case int:
				if idx := data.getIndex(k); idx >= 0 {
					data.array = append(data.array[:idx], data.array[idx+1:]...)
				}
			}
			break
		}
		//层层向下取值
		switch k := v.(type) {
		case string:
			data = data.getObject(k, false)
		case int:
			data = data.getArray(k)
		}
	}
}

// String 重构下Var的String函数,针对Set/Append后的惰性更新
func (this *Map) String(def ...string) string {
	this.refresh()
	bs, _ := this.getMarshal()(this.Var.Val())
	return string(bs)
	return this.Var.String(def...)
}

// refresh 刷新,因为是惰性加载/更新,固有个刷新函数
// 正常情况是使用不到,除非直接
func (this *Map) refresh() *Map {
	if this.hasSet {
		this.Var.Set(this.encode())
		this.hasSet = false
	}
	return this
}

/*



 */

// 获取对象的key或者数组的下标,用类型来区分,字符表示对象的key,数字表示数字的下标
// 先分割.,再正则匹配[0-9]+,例如 a.key[1][0] ,得到 {"a","key",1,0}
func (this *Map) getKeys(key string) []any {
	if len(key) == 0 {
		return []any{""}
	}
	keys := []any(nil)
	for _, v := range strings.Split(key, ".") {
		indexList := []any(nil)
		length := 0
		for _, k := range regexp.MustCompile(`\[[\-0-9]+\]`).FindAllString(v, -1) {
			length += len(k)
			//去除[]后的数字,即数组的下标
			indexList = append(indexList, Int(k[1:len(k)-1]))
		}
		if vKey := v[:len(v)-length]; len(vKey) > 0 {
			//对象的key,例 key[1][0] 中间的key
			keys = append(keys, v[:len(v)-length])
		}
		//数组的下标
		keys = append(keys, indexList...)
	}
	return keys
}

// 根据key获取对象
func (this *Map) getObject(key string, setIfNotExist bool) *Map {
	data, ok := this.object[key]
	if ok {
		return data.decode()
	}
	m := NewMap(nil, this.codec)
	if setIfNotExist {
		this.object[key] = m
	}
	return m
}

// 根据下标获取数组的元素,支持python的负数下标,
// 例如-1表示获取数组的最后一个元素
func (this *Map) getArray(idx int) *Map {
	//处理负数(支持)等情况,
	if idx = this.getIndex(idx); idx >= 0 {
		return this.array[idx].decode()
	}
	return NewMap(nil, this.codec)
}

func (this *Map) getIndex(idx int) int {
	if idx >= 0 && idx < len(this.array) {
		return idx
	}
	//对负数的支持,例如-1表示获取数组的最后一个元素,参考python
	if idx < 0 && -idx <= len(this.array) {
		return len(this.array) + idx
	}
	return -1
}

// 获取编解码函数
func (this *Map) getUnmarshal() func(data []byte, v any) error {
	if this.codec == nil {
		if codec.Default != nil {
			return codec.Default.Unmarshal
		}
		return json.Unmarshal
	}
	return this.codec.Unmarshal
}

func (this *Map) getMarshal() func(v any) ([]byte, error) {
	if this.codec == nil {
		if codec.Default != nil {
			return codec.Default.Marshal
		}
		return json.Marshal
	}
	return this.codec.Marshal
}

func (this *Map) decode() *Map {
	if !this.decoded {
		this.decoded = true
		this.object = make(map[string]*Map)
		if !this.Var.IsNil() {
			switch val := this.Var.Val().(type) {
			case map[string]any:
				for i, v := range val {
					this.object[i] = newMap(v, this.codec)
				}
			case map[any]any:
				for i, v := range val {
					this.object[String(i)] = newMap(v, this.codec)
				}
			case []any:
				this.isArray = true
				for _, v := range val {
					this.array = append(this.array, newMap(v, this.codec))
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
				//基础类型不用再次解析,字符串可以再次解析
			default:
				m := make(map[string]any)
				bs, ok := this.Var.Val().([]byte)
				if !ok {
					bs, _ = this.getMarshal()(this.Var.Val())
				}
				parse := this.getUnmarshal()
				if err := parse(bs, &m); err == nil {
					for i, v := range m {
						this.object[String(i)] = newMap(v, this.codec)
					}
				} else {
					var list []any
					if err := parse(bs, &list); err == nil {
						this.isArray = true
						for _, v := range list {
							this.array = append(this.array, newMap(v, this.codec))
						}
					}
				}

			}
		}
		this.Extend = NewExtend(this)
	}
	return this
}

func (this *Map) encode() any {

	//判断是否设置过新值,没解析直接取var的值
	if !this.hasSet {
		return this.Var.Val()
	}

	//未解析的数据,返回原始数据
	if !this.decoded {
		return this.Var.Val()
	}

	if this.isArray {
		//数组
		list := make([]any, 0)
		for _, v := range this.array {
			list = append(list, v.encode())
		}
		return list
	}

	//对象
	object := map[string]any{}
	for k, v := range this.object {
		object[k] = v.encode()
	}
	return object

}
