package main

import (
	"fmt"
	"unsafe"

	"github.com/Taguchi-ivigocon2025/workshop/unsafe/skeleton/step1/pkgA"
)

func main() {

	type B struct {
		// TODO: 公開されたint型のフィールドNを追加
		N int
	}

	var a pkgA.A
	// *A -> unsafe.Pointer -> *B
	// b := (*B)( /* TODO: *A型をunsafe.Pointerに変換する */ (&a))
	b := (*B)(unsafe.Pointer(&a))
	b.N = 100

	// 100
	fmt.Println(a.N())
}
