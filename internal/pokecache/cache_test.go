package pokecache

import (
	"testing"
	"time"
)

var duration = 5 * time.Millisecond

var items = []struct {
	key string
	val []byte
}{
	{
		key: "test key 1",
		val: []byte("Test Value 1"),
	},
	{
		key: "test key 2",
		val: []byte("Test Value 2"),
	},
}

func TestAddGet(t *testing.T) {
	cache := NewCache(duration)
	for _, item := range items {
		cache.Add(item.key, item.val)
	}
	for _, item := range items {
		test, ok := cache.Get(item.key)
		if !ok {
			t.Errorf("Key not in cache")
			return
		}
		if string(test) != string(item.val) {
			t.Errorf("Incorrect data in cache")
			return
		}
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(duration)

	cache.Add(items[0].key, items[0].val)

	_, ok := cache.Get(items[0].key)
	if !ok {
		t.Errorf("Could not find key prior to expected reaping time")
	}
	time.Sleep(duration * 3)
	_, ok = cache.Get(items[0].key)
	if ok {
		t.Errorf("Found key in cache. Should have been removed at this time.")
	}
}
