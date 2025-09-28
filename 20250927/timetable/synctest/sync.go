package synctest

import "time"

// testing.synctestの重要どころ
// 初期時刻 2000-01-01(UTC)
// 全てのゴルーチンdurably blockになったら、bubble内の時刻が進む。そしてtime.sleepは durably blockされる。
// waitはbubble内の現在のゴルーチン以外の全てのゴルーチンがdurably blockされるまでblockする。

// キャッシュのテスト
func NewCache[T any]() *Cache[T] {
	return &Cache[T]{}
}

type Cache[T any] struct {
	v       T
	ttl     time.Duration
	setTime time.Time
}

func (c *Cache[T]) Set(value T, ttl time.Duration) {
	c.v = value
	c.ttl = ttl
	c.setTime = time.Now()
}

func (c *Cache[T]) Get() T {
	if time.Since(c.setTime) >= c.ttl {
		var zero T
		return zero
	}
	return c.v
}
