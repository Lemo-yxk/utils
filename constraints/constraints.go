/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2022-03-20 01:34
**/

package constraints

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
