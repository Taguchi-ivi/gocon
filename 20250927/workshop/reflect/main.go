package main

import (
	"fmt"
	"reflect"
)

// reflect package
// 自身を鏡に反射(reflection)するようにランタイムの型や値を見ることができる
// かつそれを変更することもできる

// 補足それ以外にも
// データベースの情報を読み込んで struct にマッピングする（ORM）
// Struct Tag と go generate を使って、コード生成する
// 使い方は多いな
func main() {

	// 基本的な使い方(型の取得)
	p := Person{name: "human"}
	t := reflect.TypeOf(p)
	fmt.Println("name", t.Name()) // name Person

	// ランタイムの値を取得
	v := reflect.ValueOf(p)
	fmt.Println("value", v) // value {human}
	// fmt.Println("value", v.Elem())

}

type Person struct {
	name string
}
