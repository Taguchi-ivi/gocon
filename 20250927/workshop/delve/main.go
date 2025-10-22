package main

func calc(x, y int) int {
	a := x + x
	// b := y * y // ここがミス
	b := y + y

	return a + b
}

// dlv debug main.go

// 入力時 main.go:x(行数) breakpointを設定
// continue(cでも可) 処理を実行
// step(sでも可) 処理を1行ずつ実行
// print(pでも可) xなどで変数の値を表示
// list(lでも可) で現在の行の前後のコードを表示
// exit(qでも可) でデバッグを終了
func main() {
	var x, y int

	x = 2
	y = 3

	answer := calc(x, y)

	println(answer)
}
