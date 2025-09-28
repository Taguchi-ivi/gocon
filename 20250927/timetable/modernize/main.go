package main

import (
	"fmt"
	"sync"
)

// modernize packageのテスト
// go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...

// interface -> any
func printValue(v interface{}) {
	fmt.Println("value:", v)
}

// for i := 0; i < 10; i++ -> for range
func printLoop(max int) {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
}

// 最新の1.25まで追従している
func goroutineAndWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine is running")
	}()
	wg.Wait()
}

func main() {
	printValue("Hello, World!")

	printLoop(5)

	goroutineAndWaitGroup()
}
