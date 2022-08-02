package conv

import "time"

func GetSelect(b bool, v1, v2 interface{}) *Var {
	if b {
		return New(v1)
	}
	return New(v2)
}

func GetSelectString(b bool, v1, v2 string) string {
	return GetSelect(b, v1, v2).String()
}

func GetSelectInt(b bool, v1, v2 int) int {
	return GetSelect(b, v1, v2).Int()
}

func GetSelectInt64(b bool, v1, v2 int64) int64 {
	return GetSelect(b, v1, v2).Int64()
}

func GetSelectFloat64(b bool, v1, v2 float64) float64 {
	return GetSelect(b, v1, v2).Float64()
}

func GetSelectErr(b bool, v1, v2 error) error {
	return GetSelect(b, v1, v2).Err()
}

func GetSelectDuration(b bool, v1, v2 time.Duration) time.Duration {
	return GetSelect(b, v1, v2).Duration()
}
