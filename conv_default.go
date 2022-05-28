package conv

func GetDefault(def interface{}, param ...interface{}) *Var {
	return New(Nil.Interface(param...))
}

func GetDefaultString(def string, param ...string) string {
	return New(Nil.Interface(Interfaces(param)...)).String(def)
}

func GetDefaultInt(def int, param ...int) int {
	return New(Nil.Interface(Interfaces(param)...)).Int(def)
}

func GetDefaultInt64(def int64, param ...int64) int64 {
	return New(Nil.Interface(Interfaces(param)...)).Int64(def)
}

func GetDefaultFloat64(def float64, param ...float64) float64 {
	return New(Nil.Interface(Interfaces(param)...)).Float64(def)
}

func GetDefaultBool(def bool, param ...bool) bool {
	return New(Nil.Interface(Interfaces(param)...)).Bool(def)
}
