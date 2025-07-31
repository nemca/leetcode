package lru_cache

import (
	"testing"
)

func TestCase1(t *testing.T) {
	/*
		LRUCache lRUCache = new LRUCache(2);
		lRUCache.put(1, 1); // cache is {1=1}
		lRUCache.put(2, 2); // cache is {1=1, 2=2}
		lRUCache.get(1);    // return 1
		lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
		lRUCache.get(2);    // returns -1 (not found)
		lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
		lRUCache.get(1);    // return -1 (not found)
		lRUCache.get(3);    // return 3
		lRUCache.get(4);    // return 4
	*/
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	if v := lru.Get(1); v != 1 {
		t.Fatalf("first get(1) failed: got: %d; expected: 1\n", v)
	}
	lru.Put(3, 3)
	if v := lru.Get(2); v != -1 {
		t.Fatalf("get(2) failed: got: %d; expected: -1\n", v)
	}
	lru.Put(4, 4)
	if v := lru.Get(1); v != -1 {
		t.Fatalf("second get(1) failed: got: %d; expected: -1\n", v)
	}
	if v := lru.Get(3); v != 3 {
		t.Fatalf("get(3) failed: got: %d; expected: 3\n", v)
	}
	if v := lru.Get(4); v != 4 {
		t.Fatalf("get(4) failed: got: %d; expected: 4\n", v)
	}
}

func TestCase2(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	if v := lru.Get(1); v != -1 {
		t.Fatalf("first get(1) failed: got: %d; expected: -1\n", v)
	}
	if v := lru.Get(2); v != 3 {
		t.Fatalf("first get(1) failed: got: %d; expected: 3\n", v)
	}
}
