package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	cache.Add("https://example2.com", []byte("testdata"))
	cache.Add("https://example3.com", []byte("testdata"))
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	time.Sleep(5 * time.Second)

	cache.Add("https://example2.com", []byte("testdata"))
	cache.Add("https://example3.com", []byte("testdata"))
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
func TestCacheThreadSafety(t *testing.T) {
	cache := NewCache(5 * time.Second)

	var wg sync.WaitGroup
	numGoroutines := 100
	key := "testKey"
	expectedValue := []byte("testValue")

	// Writer Goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			cache.Add(key, expectedValue)
		}(i)
	}

	// Reader Goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			_, _ = cache.Get(key)
		}(i)
	}

	wg.Wait()

	// Check final value
	val, found := cache.Get(key)
	if !found || string(val) != string(expectedValue) {
		t.Errorf("Expected %s, got %v", expectedValue, val)
	}
}
func TestCacheReapLoopThreadSafety(t *testing.T) {
	cache := NewCache(1 * time.Second)

	var wg sync.WaitGroup
	numGoroutines := 50

	// Add initial entries
	for i := 0; i < 10; i++ {
		cache.Add(string(rune('A'+i)), []byte("type"))
	}

	// Start goroutines to access cache during cleanup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			_, _ = cache.Get("A")
		}(i)
	}

	time.Sleep(2 * time.Second) // Allow cleanup to run
	wg.Wait()

	// Verify entries have been cleaned
	for i := 0; i < 10; i++ {
		if _, found := cache.Get(string(rune('A' + i))); found {
			t.Errorf("Expected entry %c to be cleaned up", rune('A'+i))
		}
	}
}
