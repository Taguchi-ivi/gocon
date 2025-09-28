package main

import (
	"fmt"
	"sync"
)

// modernize packageのテスト
// go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./... => 指摘事項を教えてくれる
// go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -fix -test ./... => 上書き

// interface -> any
func printValue(v any) {
	fmt.Println("value:", v)
}

// for i := 0; i < 10; i++ -> for range
func printLoop(max int) {
	for i := range max {
		fmt.Println(i)
	}
}

// 最新の1.25まで追従している
func goroutineAndWaitGroup() {
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("Goroutine is running")
	})
	wg.Wait()
}

func main() {
	printValue("Hello, World!")

	printLoop(5)

	goroutineAndWaitGroup()
}
