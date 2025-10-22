package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

// 型のコピー。本来なら要素ごとに代入しないといけない
// func main() {
// 	p1 := Person{Name: "Alice", Age: 20}
// 	p2 := Person{}

// 	copyPerson(&p1, &p2)

// 	fmt.Println(p2) // {Alice 20}
// }

// // 要素ごとのコピーは要素が増えるとここら辺も対応する必要がある。
// func copyPerson(src, dst *Person) {
// 	dst.Name = src.Name
// 	dst.Age = src.Age
// }

// reflectを使うと
func main() {
	p1 := Person{Name: "Alice", Age: 20}
	p2 := Person{}

	copyPerson(&p1, &p2)

	fmt.Println(p2) // {Alice 20}
}

func copyPerson(src, dst *Person) {
	// reflectを使うと要素数が増えても対応できる
	// `reflect.ValueOf()` と `Value.Elem()` を使うと、ポインター型の値情報から値そのものを取り出せます
	srcValue := reflect.ValueOf(src).Elem()
	fmt.Println("srcValue:", srcValue)
	dstValue := reflect.ValueOf(dst).Elem()
	fmt.Println("dstValue:", dstValue)

	for i := 0; i < srcValue.NumField(); i++ {
		dstValue.Field(i).Set(srcValue.Field(i))
	}
}
