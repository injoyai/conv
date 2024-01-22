package conv

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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
		return value.Unix()
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
		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			if v, err := strconv.ParseInt(s[2:], 16, 64); err == nil {
				return v * base
			}
		}
		// BIN 二进制
		if len(s) > 2 && s[0] == '0' && (s[1] == 'b' || s[1] == 'B') {
			if v, err := strconv.ParseInt(s[2:], 2, 64); err == nil {
				return v * base
			}
		}
		// DEC 十进制
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			return v * base
		}
		// 时间戳
		if d, err := time.ParseDuration(s); err == nil {
			return int64(d)
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
