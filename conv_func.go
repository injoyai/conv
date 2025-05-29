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

// Range 限制范围,设置默认值,组合函数,conv.Range(1,1,10,2)
func Range[T Comparable](_default T, min, max T, vals ...T) T {
	val := Default(_default, vals...)
	if min > max {
		min, max = max, min
	}
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
