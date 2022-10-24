package conv

import "time"

func Select(b bool, v1, v2 interface{}) interface{} {
	if b {
		return v1
	}
	return v2
}

func SelectVar(b bool, v1, v2 interface{}) *Var {
	return New(Select(b, v1, v2))
}

func SelectString(b bool, v1, v2 string) string {
	return SelectVar(b, v1, v2).String()
}

// GetSelectString 过度
func GetSelectString(b bool, v1, v2 string) string {
	return SelectVar(b, v1, v2).String()
}

func SelectUint8(b bool, v1, v2 uint8) uint8 {
	return SelectVar(b, v1, v2).Uint8()
}

func SelectInt(b bool, v1, v2 int) int {
	return SelectVar(b, v1, v2).Int()
}

func SelectInt64(b bool, v1, v2 int64) int64 {
	return SelectVar(b, v1, v2).Int64()
}

func SelectUint(b bool, v1, v2 uint) uint {
	return SelectVar(b, v1, v2).Uint()
}

func SelectUint64(b bool, v1, v2 uint64) uint64 {
	return SelectVar(b, v1, v2).Uint64()
}

func SelectFloat64(b bool, v1, v2 float64) float64 {
	return SelectVar(b, v1, v2).Float64()
}

func SelectErr(b bool, v1, v2 error) error {
	return SelectVar(b, v1, v2).Err()
}

func SelectDuration(b bool, v1, v2 time.Duration) time.Duration {
	return SelectVar(b, v1, v2).Duration()
}
