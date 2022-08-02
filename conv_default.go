package conv

import "time"

func GetDefault(def interface{}, param ...interface{}) *Var {
	return New(New(nil).Interface(param...))
}

func GetDefaultString(def string, param ...string) string {
	return New(New(nil).Interface(Interfaces(param)...)).String(def)
}

func GetDefaultUint(def uint, param ...uint) uint {
	return New(New(nil).Interface(Interfaces(param)...)).Uint(def)
}

func GetDefaultInt(def int, param ...int) int {
	return New(New(nil).Interface(Interfaces(param)...)).Int(def)
}

func GetDefaultInt64(def int64, param ...int64) int64 {
	return New(New(nil).Interface(Interfaces(param)...)).Int64(def)
}

func GetDefaultFloat64(def float64, param ...float64) float64 {
	return New(New(nil).Interface(Interfaces(param)...)).Float64(def)
}

func GetDefaultBool(def bool, param ...bool) bool {
	return New(New(nil).Interface(Interfaces(param)...)).Bool(def)
}

func GetDefaultErr(def error, param ...error) error {
	return New(New(nil).Interface(Interfaces(param)...)).Err(def)
}

func GetDefaultDuration(def time.Duration, param ...time.Duration) time.Duration {
	return New(New(nil).Interface(Interfaces(param)...)).Duration(int(def))
}
