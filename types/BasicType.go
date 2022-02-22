package types

import (
	"fmt"
	"unsafe"
)

func BasicType() {
	var s string
	s = "Hello World"
	fmt.Printf("%s byte \n", s)

	var c_1 complex64
	fmt.Printf("The size of complex64 is %d byte \n", unsafe.Sizeof(c_1))

	var c_2 complex128
	fmt.Printf("The size of complex128 is %d byte \n", unsafe.Sizeof(c_2))

	var f_1 float32
	fmt.Printf("The size of float32 is %d byte \n", unsafe.Sizeof(f_1))

	var f_2 float64
	fmt.Printf("The size of float64 is %d byte \n", unsafe.Sizeof(f_2))

	var ui_8 uint8
	fmt.Printf("The size of uint8 is %d byte \n", unsafe.Sizeof(ui_8))

	var ui_16 uint16
	fmt.Printf("The size of uint16 is %d byte \n", unsafe.Sizeof(ui_16))

	var ui_32 uint32
	fmt.Printf("The size of uint32 is %d byte \n", unsafe.Sizeof(ui_32))

	var ui_64 uint64
	fmt.Printf("The size of uint64 is %d byte \n", unsafe.Sizeof(ui_64))

	var i_8 int8
	fmt.Printf("The size of int8 is %d byte \n", unsafe.Sizeof(i_8))

	var i_16 int16
	fmt.Printf("The size of int16 is %d byte \n", unsafe.Sizeof(i_16))

	var i_32 int32
	fmt.Printf("The size of int32 is %d byte \n", unsafe.Sizeof(i_32))

	var i_64 int64
	fmt.Printf("The size of int64 is %d byte \n", unsafe.Sizeof(i_64))

	var ru rune
	fmt.Printf("The size of rune is %d byte \n", unsafe.Sizeof(ru))

	var by byte
	fmt.Printf("The size of byte is %d byte \n", unsafe.Sizeof(by))

	var u_p uintptr
	fmt.Printf("The size of uintptr is %d byte \n", unsafe.Sizeof(u_p))

	var err error
	fmt.Printf("The size of error is %d byte \n", unsafe.Sizeof(err))

	var b bool
	fmt.Printf("The size of bool is %d byte \n", unsafe.Sizeof(b))
}
