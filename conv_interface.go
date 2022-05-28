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

type apiInterfaces interface {
	Interfaces() []interface{}
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
