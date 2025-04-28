package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type KVStore struct {
	data map[string]item
	mu   sync.RWMutex
}

type item struct {
	value      any
	expiration int64 // Unix timestamp (seconds)
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]item),
	}
}

func (s *KVStore) Set(key string, value any, ttl time.Duration) {
	s.mu.Lock()
	s.data[key] = item{value: value, expiration: time.Now().Add(ttl).Unix()}
	s.mu.Unlock()
}

func (s *KVStore) Get(key string) (any, bool) {

	s.mu.RLock()
	itemFound, exists := s.data[key]
	if !exists {
		return item{}, false
	}

	if itemFound.expiration < time.Now().Unix() {
		return item{}, false
	}
	s.mu.RUnlock()

	return itemFound, true
}

func (s *KVStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)

}

func (s *KVStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	count := 0
	for _, value := range s.data {
		if value.expiration > time.Now().Unix() {
			count++
		}
	}
	return count
}

func (s *KVStore) StartGC(interval time.Duration) {

	go func() {

		ticker := time.Tick(interval)

		for {
			select {
			case <-ticker:
				for key, value := range s.data {
					if value.expiration < time.Now().Unix() {
						delete(s.data, key)
					}
				}
			}
		}

	}()
}

func main() {
	kv := NewKVStore()

	var wg sync.WaitGroup
	kv.StartGC(time.Second)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			kv.Set(strconv.Itoa(i), i+1, time.Duration(i))
			wg.Done()
		}()
	}

	wg.Wait()
	value, ok := kv.Get(strconv.Itoa(8))
	if !ok {
		fmt.Println("Error fetching item")
		return
	}

	itemValue := value.(item)

	fmt.Println(itemValue.value)

}
