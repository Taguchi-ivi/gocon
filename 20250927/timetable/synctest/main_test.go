package synctest

import (
	"fmt"
	"testing"
	"testing/synctest"
	"time"
)

// 参考: https://zenn.dev/canary_techblog/articles/ec8a96b4541685
// go docを読もう: GOEXPERIMENT=synctest go doc testing/synctest.Run

// GOEXPERIMENT=synctest go test -run "TestCache_Get" // synctest.からスタートしないと5秒以上かかる
// go test ./... // 5秒以上かかる

func TestCache_Set(t *testing.T) {
	ttl := 5 * time.Second
	cache := NewCache[string]()
	cache.Set("cached item", ttl)

	if got := cache.Get(); got != "cached item" {
		t.Errorf("expected 'cached item'; got %v", got)
	}

	time.Sleep(ttl)

	if got := cache.Get(); got != "" {
		t.Errorf("expected ''; got %v", got)
	}
}

// GOEXPERIMENT=synctest go test -run "TestCache_Get" これで速攻で終わるようになる
func TestCache_Get(t *testing.T) {
	synctest.Run(func() {
		fmt.Println(time.Now()) // 現在時刻を出力
		ttl := 5 * time.Second
		cache := NewCache[string]()
		cache.Set("cached item", ttl)

		if got := cache.Get(); got != "cached item" {
			t.Errorf("expected 'cached item'; got %v", got)
		}

		fmt.Println(time.Now()) // 現在時刻を出力
		time.Sleep(ttl)
		fmt.Println(time.Now()) // 現在時刻を出力, ここから時刻が進む

		if got := cache.Get(); got != "" {
			t.Errorf("expected ''; got %v", got)
		}
		fmt.Println(time.Now()) // 現在時刻を出力
	})
}
