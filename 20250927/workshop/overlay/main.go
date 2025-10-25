package main

// Goのビルドでは、overlayオプションを使うことで特定のファイルを別のファイルで置き換えてビルドできる。

// 現在時刻を表示するプログラムの現在を固定させるためにどうやってoverlayを使うか学ぶ
// overlay.json内の値を go env GOROOTコマンドで取得できる値に置き換える(下記のGOROOT部分)
// "${GOROOT}/src/time/time.go": "./overlay/example/time/time.go"

// go run -overlay ./overlay/example/overlay.json ./overlay/example

// go buildや go testでは -overlayオプションを提供している。

// そのままgo testしてもうまくいかないが-overlayで置き換えてテストするとうまくいく
// go test -overlay ./overlay/example/overlay.json ./overlay/example
