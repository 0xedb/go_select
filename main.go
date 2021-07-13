package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 1
	b := 3

	ad := uintptr(unsafe.Pointer(&a))
	bd := uintptr(unsafe.Pointer(&b))
	
	fmt.Println(ad ^ bd)
	fmt.Println(bd ^ 0)
}