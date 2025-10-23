// sample.go - toolexec の動作確認用プログラム
package main

import (
	"fmt"
)

// ビルド時に -ldflags で値を設定できる変数
var (
	buildTime    = "unknown"
	buildVersion = "dev"
)

func main() {
	fmt.Println(sayFromGoConference())
}

func sayFromGoConference() string {
	return "Hello, Gopher Welcome to Go Conference 2025!"
}

/*
== 実践的な使い方 ==
各パッケージのコンパイル時間を記録
デバッグ情報の自動除去
ビルド情報の自動埋め込み
*/
