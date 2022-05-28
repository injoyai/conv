package conv

import (
	"reflect"
	"time"
)

// IsDefault 是否是默认值
func IsDefault(i interface{}) bool {
	return IsZero(i)
}

// IsZero 是否是默认值
func IsZero(i interface{}) bool {
	if i == nil {
		return true
	}
	r := reflect.ValueOf(GetNature(i))
	if !r.CanInterface() {
		return true
	}
	switch r.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Slice, reflect.Chan, reflect.Map:
		return r.Len() == 0
	default:
		return r.IsZero()
	}
}

// IsNumber 是否是数字类型
func IsNumber(i interface{}) bool {
	return IsInt(i) || IsFloat(i)
}

// IsString 是否是字符串
func IsString(i interface{}) bool {
	_, ok := GetNature(i).(string)
	return ok
}

// IsInt 是否是整数类型
func IsInt(i interface{}) bool {
	switch GetNature(i).(type) {
	case int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64,
		float32,
		float64:
		return true
	}
	return false
}

// IsFloat 是否是浮点类型
func IsFloat(i interface{}) bool {
	switch GetNature(i).(type) {
	case float32,
		float64:
		return true
	}
	return false
}

// IsBool 是否是布尔类型
func IsBool(i interface{}) bool {
	_, ok := GetNature(i).(bool)
	return ok
}

// IsTime 是否是时间类型
func IsTime(i interface{}) bool {
	_, ok := GetNature(i).(time.Time)
	return ok
}

// IsArray 是否是数组
func IsArray(i interface{}) bool {
	switch GetNatureKind(i) {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

// IsPointer 是否是指针
func IsPointer(i interface{}) bool {
	return GetNatureKind(i) == reflect.Ptr
}

// IsNil 是否是nil
func IsNil(i interface{}) bool {
	return i == nil
}

// GetNatureKind 获取kind(无指针)
func GetNatureKind(i interface{}) reflect.Kind {
	return reflect.ValueOf(GetNature(i)).Kind()
}

// GetNature 去除指针
func GetNature(i interface{}) interface{} {
	if i != nil {
		r := reflect.ValueOf(i)
		kind := r.Kind()
		for kind == reflect.Ptr {
			r = r.Elem()
			kind = r.Kind()
		}
		if r.IsValid() && r.CanInterface() {
			return r.Interface()
		}
	}
	return nil
}
