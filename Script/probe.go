package main

import (
	"github.com/GeeCache/geecache/lru"
	"log"
)
type String string

func (d String) Len() int {
	return len(d)
}

func TestRemoveoldest() {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lruCache := lru.New(int64(cap), nil)
	lruCache.Add(k1, String(v1))
	lruCache.Add(k2, String(v2))
	lruCache.Add(k3, String(v3))

	if _, ok := lruCache.Get("key1"); ok || lruCache.Len() != 2 {
		log.Printf("Removeoldest key1 failed")
	}
}

func TestGet() {
	// 这个地方把 maxBytes 设为 0，表示这个 cache 的容量是无穷大
	lruCache := lru.New(0, nil)
	lruCache.Add("key1", String("1234"))
	if v, ok := lruCache.Get("key1"); !ok || string(v.(String)) != "1234" {
		log.Printf("cache hit key1=1234 failed")
	}

	if _, ok := lruCache.Get("key1"); !ok {
		log.Print("cache miss key2 failed")
	}
}

func main() {
	TestRemoveoldest()
}
