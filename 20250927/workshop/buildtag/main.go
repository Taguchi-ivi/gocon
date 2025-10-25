package main

// ビルドタグについて学ぶ
// ビルドタグを使った攻撃方法について学ぶ
// ビルドタグを使った攻撃への対応方法について学ぶ

// ビルドタグは特定の条件下でのみコードをコンパイルするための仕組みです。(ex. //go:build go1.25)
// これによって 現在は無害だが、将来のGoバージョンで悪意のあるコードが実行される可能性もある
// そもそもビルドされないため、テストで発見しにくい。

// gotoolchainを使って、vを変更する GOTOOLCHAIN=go1.24.6 go run ./skeleton/step1
// ビルドタグの検出チェックをするならgostaticanalysis/buildtagツールを使う
// もしくはgo/build/constraintパッケージ
