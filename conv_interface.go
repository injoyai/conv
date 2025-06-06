package conv

type apiString interface {
	String() string
}

type apiError interface {
	Error() string
}

type apiGoString interface {
	GoString() string
}

// mqtt
type apiPayload interface {
	Payload() string
}

type apiInterfaces interface {
	Interfaces() []any
}

type apiFloat32 interface {
	Float32() float32
}

type apiFloat64 interface {
	Float64() float64
}

type apiInt interface {
	Int() int
}

type apiInt64 interface {
	Int64() int64
}

type apiUint interface {
	Uint() uint
}

type apiUint64 interface {
	Uint64() uint64
}

type apiBool interface {
	Bool() bool
}

type apiCopy interface {
	Copy() any
}

// Comparable 可比较,同1.21版本的cmp.Ordered
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Floater interface {
	~float32 | ~float64
}
