package conv

// Default 默认值,conv.Default(1,2)
func Default[T any](_default T, param ...T) T {
	if len(param) > 0 {
		return param[0]
	}
	return _default
}

// Select 模拟三元表达式,conv.Select(true,1,2)
func Select[T any](b bool, v1, v2 T) T {
	if b {
		return v1
	}
	return v2
}
