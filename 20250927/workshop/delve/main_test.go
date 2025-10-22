package main

import "testing"

// delveはtestでも使える。かつdelveはライブラリにも使える。test関数で呼べばいいというね。とてもいい！！！
// dlv test ./ (こちらはpackageのテストを実行できる)
// dlv test ./ -- -test.run Test_calc (こちらは関数名を指定して実行できる)
// break main_test.go:10 (こちらはブレークポイントを設定できる)
// 特定の関数やメソッドに対応できるのでいい！！
func Test_calc(t *testing.T) {
	t.Parallel()

	want := 10
	got := calc(2, 3)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
