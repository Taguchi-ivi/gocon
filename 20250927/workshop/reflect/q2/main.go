package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	DNA  string
	Soul string `copyable:"false"`
}

type Food struct {
	Name             string
	Kind             string
	secretIngredient string
}

func main() {
	p1 := Person{Name: "Alice", DNA: "ALICE_DNA", Soul: "ALICE_SOUL"}
	p2 := Person{}
	f1 := Food{Name: "Icecream", Kind: "Sweet", secretIngredient: "Salt"}
	f2 := Food{}

	copyStruct(&p1, &p2)
	copyStruct(&f1, &f2)

	fmt.Println(p2) // {Alice 20}
	fmt.Println(f2) // {Icecream Sweet}
}

// TODO: implement
// 任意の型の struct をコピーできる copyStruct 関数を実装してみましょう。 雛形のコードはこちらです。任意の型の struct をコピーできる copyStruct 関数を実装してみましょう。 雛形のコードはこちらです。
// `any` を使うと、任意の型を引数で受け取れるようになります
func copyStruct(src, dst any) {
	// reflectを使うと要素数が増えても対応できる
	// `reflect.ValueOf()` と `Value.Elem()` を使うと、ポインター型の値情報から値そのものを取り出せます
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		dstValue.Field(i).Set(srcValue.Field(i))
	}
}
