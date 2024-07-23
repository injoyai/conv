package conv

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const IntSize = strconv.IntSize

var (
	emptyMap = map[string]*struct{}{
		"":      {},
		"0":     {},
		"no":    {},
		"off":   {},
		"false": {},
		"close": {},
		"关":     {},
		"否":     {},
		"假":     {},
	}
	octMap = map[uint64]string{
		0: "0", 1: "1", 2: "2", 3: "3",
		4: "4", 5: "5", 6: "6", 7: "7",
	}
)

// toBytes 任意类型转 []byte.
func toBytes(i interface{}) []byte {
	if i == nil {
		return []byte{}
	}
	switch value := GetNature(i).(type) {
	case []byte:
		return value
	case []bool:
		result := byte(0)
		for _, b := range value {
			result *= 2
			if b {
				result++
			}
		}
		return toBytes(result)
	case *Var:
		return value.Bytes()
	}
	if IsNumber(i) {
		// int 类型无法解析
		if val, ok := i.(int); ok {
			switch IntSize {
			case 32:
				i = int32(val)
			default:
				i = int64(val)
			}
		}
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, i)
		value := bytesBuffer.Bytes()
		return value
	}
	return []byte(toString(i))
}

// toString 任意类型转 string.
func toString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch value := i.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		return value.String()
	case time.Duration:
		return value.String()
	case apiString:
		return value.String()
	case apiError:
		return value.Error()
	case apiGoString:
		return value.GoString()
	case apiPayload:
		return value.Payload()
	case io.Reader:
		bs, _ := ioutil.ReadAll(value)
		return string(bs)
	case *Var:
		return value.String()
	default:
		if value == nil {
			return ""
		}
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		}
		if kind == reflect.Ptr {
			return toString(rv.Elem().Interface())
		}
		// we try use json.Marshal to convert.
		jsonBytes, err := json.Marshal(value)
		if err == nil {
			return string(jsonBytes)
		}
		return fmt.Sprint(value)

	}
}

// toInt64 任意类型转 int64.
func toInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		//长度如果超过8位,会舍弃8位之后的字节
		//正序 大端 []byte{0x00,0x00,0x01} >>> 1
		return int64(binary.BigEndian.Uint64(padding(value, 8)))
	case []bool:
		// BIN 二进制
		result := int64(0)
		for _, b := range value {
			result *= 2
			if b {
				result++
			}
		}
		return result
	case time.Time:
		return value.UnixNano()
	case time.Duration:
		return int64(value)
	case apiInt:
		return int64(value.Int())
	case apiInt64:
		return value.Int64()
	case *Var:
		return value.Int64()
	default:
		s := toString(value)
		base := int64(1)
		if len(s) > 0 {
			if s[0] == '-' {
				base = -1
				s = s[1:]
			} else if s[0] == '+' {
				s = s[1:]
			}
		}
		// HEX 十六进制
		if len(s) > 2 && strings.ToLower(s[0:2]) == "0x" {
			if v, err := strconv.ParseInt(s[2:], 16, 64); err == nil {
				return v * base
			}
		}
		// BIN 二进制
		if len(s) > 2 && strings.ToLower(s[0:2]) == "0b" {
			if v, err := strconv.ParseInt(s[2:], 2, 64); err == nil {
				return v * base
			}
		}
		// DEC 十进制
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			return v * base
		}
		// 时间
		if d, err := time.ParseDuration(s); err == nil {
			return int64(d) * base
		}
		// Float64
		return int64(toFloat64(value))
	}
}

// toUint64 任意类型转 uint64.
func toUint64(i interface{}) uint64 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case bool:
		if value {
			return 1
		}
		return 0
	case []byte:
		return binary.BigEndian.Uint64(padding(value, 8))
	case []bool:
		// BIN 二进制
		result := uint64(0)
		for _, b := range value {
			result *= 2
			if b {
				result++
			}
		}
		return result
	case time.Time:
		return uint64(value.UnixNano())
	case time.Duration:
		return uint64(value)
	case apiUint:
		return uint64(value.Uint())
	case apiUint64:
		return value.Uint64()
	case *Var:
		return value.Uint64()
	default:
		s := toString(value)
		// HEX 十六进制
		if len(s) > 2 && strings.ToLower(s[0:2]) == "0x" {
			if v, err := strconv.ParseUint(s[2:], 16, 64); err == nil {
				return v
			}
		}
		// BIN 二进制
		if len(s) > 2 && strings.ToLower(s[0:2]) == "0b" {
			if v, err := strconv.ParseUint(s[2:], 2, 64); err == nil {
				return v
			}
		}
		// DEC 十进制
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			return v
		}
		// 时间戳
		if d, err := time.ParseDuration(s); err == nil {
			return uint64(d)
		}
		// Float64
		return uint64(toFloat64(value))
	}
}

// toFloat64 任意类型转 float64.
func toFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case float32:
		// 处理互转精度问题,实际内存字节数据变了,需要注意
		return toFloat64(toString(value))
	case float64:
		return value
	case uint32: //IEEE754标准
		return float64(math.Float32frombits(value))
	case uint64: //IEEE754标准
		return math.Float64frombits(value)
	case apiFloat32:
		return float64(value.Float32())
	case apiFloat64:
		return value.Float64()
	case []byte:
		if len(value) <= 4 {
			// 处理互转精度问题
			return toFloat64(math.Float32frombits(binary.BigEndian.Uint32(padding(value, 4))))
		}
		return math.Float64frombits(binary.BigEndian.Uint64(padding(value, 8)))
	case *Var:
		return value.Float64()
	default:
		v, _ := strconv.ParseFloat(toString(i), 64)
		return v
	}
}

// toBool 任意类型转 bool.
func toBool(i interface{}) bool {
	if i == nil {
		return false
	}
	switch value := i.(type) {
	case bool:
		return value
	case []byte, string:
		return emptyMap[strings.ToLower(toString(value))] == nil
	case apiBool:
		return value.Bool()
	case *Var:
		return value.Bool()
	default:
		rv := reflect.ValueOf(i)
		switch rv.Kind() {
		case reflect.Ptr:
			return !rv.IsNil()
		case reflect.Map:
			fallthrough
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			return rv.Len() != 0
		case reflect.Struct:
			return true
		default:
			return emptyMap[strings.ToLower(toString(i))] == nil
		}
	}
}

// toInterfaces 任意类型转 []interface{}.
func toInterfaces(i interface{}) []interface{} {
	if i == nil {
		return nil
	}
	if r, ok := i.([]interface{}); ok {
		return r
	}
	var array []interface{}
	switch value := i.(type) {
	case []string:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int8:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int16:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int32:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []int64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint8:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint16:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []uint32:
		for _, v := range value {
			array = append(array, v)
		}
	case []uint64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []bool:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []float32:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case []float64:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v
		}
	case apiInterfaces:
		return value.Interfaces()
	case *Var:
		return []interface{}{value.Val()}
	case []*Var:
		array = make([]interface{}, len(value))
		for k, v := range value {
			array[k] = v.Val()
		}
	default:
		var (
			rv   = reflect.ValueOf(i)
			kind = rv.Kind()
		)
		for kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Slice, reflect.Array:
			array = make([]interface{}, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				array[i] = rv.Index(i).Interface()
			}
		default:
			return []interface{}{i}
		}
	}
	return array
}

func toGMap(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	_ = json.Unmarshal(Bytes(i), &m)
	return m
}

func toIMap(i interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	_ = json.Unmarshal(Bytes(i), &m)
	return m
}

func toSMap(i interface{}) map[string]string {
	m := make(map[string]string)
	_ = json.Unmarshal(Bytes(i), &m)
	return m
}

// unmarshal 任意类型解析到ptr
// 切片类型会覆盖原有数据,字典和结构类型会保留原有字段,同json解析
func unmarshal(i, ptr interface{}, param ...UnmarshalParam) error {
	if i == nil || ptr == nil {
		return nil
	}

	valValue := reflect.ValueOf(i)
	valType := reflect.TypeOf(i)
	ptrValue := reflect.ValueOf(ptr)
	ptrType := reflect.TypeOf(ptr)

	if ptrType.Kind() != reflect.Ptr {
		return errors.New("参数(ptr)需要指针类型")
	}

	for ptrType.Kind() == reflect.Ptr {
		if ptrValue.IsNil() {
			//对为分配内存的类型,进行内存分配
			ptrValue.Set(reflect.New(ptrType.Elem()))
		}
		ptrValue = ptrValue.Elem()
		ptrType = ptrType.Elem()
	}

	for valType.Kind() == reflect.Ptr {
		valValue = valValue.Elem()
		valType = valType.Elem()
	}

	//判断入参是否是字符或者字节数组,并且指针是对象,字典和切片,则使用json解析
	//入参是 "s",*[]string{} >>> *[]string{"s"} 推荐使用Strings实现
	switch ptrType.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		switch val := valValue.Interface().(type) {
		case string, []byte:
			//todo is a good idea?
			if bs := Bytes(val); len(bs) >= 2 &&
				(bs[0] == '[' || bs[0] == '{') &&
				(bs[len(bs)-1] == ']' || bs[len(bs)-1] == '}') {
				return json.Unmarshal(Bytes(val), ptr)
			}
		}
	}

	switch ptrType.Kind() {

	case reflect.Struct:
		return copyStruct(ptrValue, valValue, param...)

	case reflect.Map:
		return copyMap(ptrValue, valValue, param...)

	case reflect.Slice:
		return copySlice(ptrValue, valValue)

	default:
		//基础类型的复制
		copyBaseKind(ptrValue, valValue)

	}

	return nil
}

func copyMap(result reflect.Value, original reflect.Value, params ...UnmarshalParam) error {
	if result.Kind() != reflect.Map {
		return nil
	}

	//声明内存空间
	if result.IsNil() {
		result.Set(reflect.MakeMap(result.Type()))
	}

	switch original.Kind() {
	case reflect.Map:
		x := original.MapRange()
		m := make(map[reflect.Value]reflect.Value)
		for x.Next() {
			m[x.Key()] = x.Value()
		}
		for k, v := range m {
			val := reflect.New(result.Type().Elem()).Elem()
			if err := copyValue(val, v); err != nil {
				return err
			}
			result.SetMapIndex(k, val)
		}

	case reflect.Struct:
		for idx := 0; idx < original.NumField(); idx++ {
			f := original.Type().Field(idx)
			val := reflect.New(result.Type().Elem()).Elem()
			if err := copyValue(val, original.Field(idx)); err != nil {
				return err
			}

			// 解析对象,尝试映射tag上,默认用字段的tag(json)
			tags := []string{"json"}
			if len(params) > 0 {
				tags = params[0].Tags
			}

			var key string
			var hasTag bool
			for _, tag := range tags {
				if key, hasTag = f.Tag.Lookup(tag); hasTag {
					break
				}
			}
			if !hasTag {
				key = f.Name
			}
			//用对象的字段名称
			result.SetMapIndex(reflect.ValueOf(key), val)
		}

	}

	return nil
}

func copyStruct(result reflect.Value, original reflect.Value, param ...UnmarshalParam) error {
	if result.Kind() != reflect.Struct {
		return nil
	}

	//判断是否分配内存空间,并分配
	if reflect.DeepEqual(result.Interface(), reflect.Zero(result.Type()).Interface()) {
		result.Set(reflect.New(result.Type()).Elem())
	}

	//尝试通过tag映射的key,默认tag的key为json
	tags := []string{"json"}
	if len(param) > 0 {
		tags = param[0].Tags
	}

	switch original.Kind() {
	case reflect.Struct:
		fieldMap := make(map[string]reflect.Value)
		tagMaps := make([]map[string]reflect.Value, len(tags))
		for idx := 0; idx < original.NumField(); idx++ {
			//保存字段对应的值
			fieldMap[original.Type().Field(idx).Name] = original.Field(idx)
			//查看原始对象数据的字段是否存在预期的tag,存在则保存tag和对应的字段值
			for tagIndex, tag := range tags {
				if key, ok := original.Type().Field(idx).Tag.Lookup(tag); ok {
					if tagMaps[tagIndex] == nil {
						tagMaps[tagIndex] = make(map[string]reflect.Value)
					}
					tagMaps[tagIndex][key] = original.Field(idx)
				}
			}
		}
		for idx := 0; idx < result.NumField(); idx++ {
			if field := result.Field(idx); field.CanSet() {
				var val reflect.Value
				var hasTagValue bool
				//获取写入对象的tag,遍历tag,是否有预期的tag,例如json
				for tagIndex, tag := range tags {
					if key, hasTag := result.Type().Field(idx).Tag.Lookup(tag); hasTag {
						if val, hasTagValue = tagMaps[tagIndex][key]; hasTagValue {
							break
						}
					}
				}
				//从tag的key获取不到对应的值,再尝试通过对象结构的字段名
				if !hasTagValue {
					val = fieldMap[result.Type().Field(idx).Name]
				}
				if err := copyValue(field, val); err != nil {
					return err
				}
			}
		}

	case reflect.Map:
		x := original.MapRange()
		m := make(map[string]reflect.Value)
		for x.Next() {
			m[String(x.Key().Interface())] = x.Value()
		}

		for idx := 0; idx < result.NumField(); idx++ {
			if field := result.Field(idx); field.CanSet() {
				var val reflect.Value
				var hasTagValue bool
				//先根据tag的key(例json),尝试获取map的值
				for _, tag := range tags {
					if key, hasTag := result.Type().Field(idx).Tag.Lookup(tag); hasTag {
						if val, hasTagValue = m[key]; hasTagValue {
							break
						}
					}
				}
				//根据tag的key获取不到map的值,再尝试通过对象结构的字段名
				if !hasTagValue {
					val = m[result.Type().Field(idx).Name]
				}
				if err := copyValue(field, val); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// copySlice 会覆盖原有数据,同json
func copySlice(result, original reflect.Value) error {
	if result.Kind() != reflect.Slice {
		return nil
	}

	switch original.Kind() {
	case reflect.Slice:
		result.Set(reflect.MakeSlice(result.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			if err := copyValue(result.Index(i), original.Index(i)); err != nil {
				return err
			}
		}

	default:
		result.Set(reflect.MakeSlice(result.Type(), 1, 1))
		if err := copyValue(result.Index(0), original); err != nil {
			return err
		}

	}

	return nil
}

func copyValue(result, original reflect.Value) error {
	if !result.CanAddr() {
		return nil
	}
	if !original.IsValid() {
		return nil
	}
	if !original.CanInterface() {
		return nil
	}
	return unmarshal(original.Interface(), result.Addr().Interface())
}

// copyBaseKind 基础类型的复制
func copyBaseKind(result, original reflect.Value) {
	if result.Kind() == original.Kind() {
		if result.CanSet() {
			result.Set(original)
		}
		return
	}
	switch result.Kind() {
	case reflect.Interface:
		//这里注意,会改变result的类型,为original的类型
		if result.CanSet() {
			result.Set(original)
		}

	case reflect.String:
		copyBaseKind(result, reflect.ValueOf(String(original.Interface())))

	case reflect.Bool:
		copyBaseKind(result, reflect.ValueOf(Bool(original.Interface())))

	case reflect.Int:
		copyBaseKind(result, reflect.ValueOf(Int(original.Interface())))

	case reflect.Int8:
		copyBaseKind(result, reflect.ValueOf(Int8(original.Interface())))

	case reflect.Int16:
		copyBaseKind(result, reflect.ValueOf(Int16(original.Interface())))

	case reflect.Int32:
		copyBaseKind(result, reflect.ValueOf(Int32(original.Interface())))

	case reflect.Int64:
		copyBaseKind(result, reflect.ValueOf(Int64(original.Interface())))

	case reflect.Uint:
		copyBaseKind(result, reflect.ValueOf(Uint(original.Interface())))

	case reflect.Uint8:
		copyBaseKind(result, reflect.ValueOf(Uint8(original.Interface())))

	case reflect.Uint16:
		copyBaseKind(result, reflect.ValueOf(Uint16(original.Interface())))

	case reflect.Uint32:
		copyBaseKind(result, reflect.ValueOf(Uint32(original.Interface())))

	case reflect.Uint64:
		copyBaseKind(result, reflect.ValueOf(Uint64(original.Interface())))

	case reflect.Float32:
		copyBaseKind(result, reflect.ValueOf(Float32(original.Interface())))

	case reflect.Float64:
		copyBaseKind(result, reflect.ValueOf(Float64(original.Interface())))

	case reflect.Complex64, reflect.Complex128,
		reflect.Chan, reflect.Func,
		reflect.Invalid, reflect.UnsafePointer,
		reflect.Struct, reflect.Map, reflect.Slice:
		//这些类型暂时还不支持
	}
}

// copySameKind 递归复制值，处理指针的情况
func copySameKind(result, original reflect.Value) {
	if result.Kind() != original.Kind() {
		return
	}

	switch original.Kind() {
	case reflect.Ptr:
		// 如果是指针，则递归复制指针指向的内容
		// result的值暂时还是nil,无需判断
		if !original.IsNil() {
			result.Set(reflect.New(original.Type().Elem()))
			copySameKind(result.Elem(), original.Elem())
		}

	case reflect.Struct:
		// 如果是结构体，则递归复制结构体的字段
		for i := 0; i < original.NumField(); i++ {
			if result.Field(i).CanSet() {
				copySameKind(result.Field(i), original.Field(i))
			}
		}

	case reflect.Map:
		if !original.IsNil() {
			// 如果是映射，则创建新的映射并递归复制键值对
			result.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				destKey := reflect.New(key.Type()).Elem()
				copySameKind(destKey, key)
				destValue := reflect.New(original.MapIndex(key).Type()).Elem()
				copySameKind(destValue, original.MapIndex(key))
				result.SetMapIndex(destKey, destValue)
			}
		}

	case reflect.Slice:
		if !original.IsNil() {
			// 如果是切片，则创建新的切片并递归复制元素
			result.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
			for i := 0; i < original.Len(); i++ {
				copySameKind(result.Index(i), original.Index(i))
			}
		}

	default:

		result.Set(original)

	}
}

type UnmarshalParam struct {
	Tags []string //自定义解析tag,默认tag为json
}
