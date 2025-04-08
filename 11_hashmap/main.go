package main

import (
	"errors"
	"fmt"
	"reflect"
)

// node for storing key value pair
type node[K comparable, V comparable] struct {
	key   K
	value V
	next  *node[K, V]
}

// hashmap for storing multiple key value pairs
type hashmap[K comparable, V comparable] struct {
	capacity     int
	size         int
	arrHead      []*node[K, V]
	maxThreshold float32
}

// function for generating custom hash
func (h *hashmap[K, V]) generateHash(key K) int {
	keyType := reflect.TypeOf(key)

	if keyType.Kind() == reflect.Int {
		x := reflect.ValueOf(key)
		return int(x.Int())
	}
	return 0
}

// function for hashing key
func (h *hashmap[K, V]) hashFn(v int) int {
	return v % h.capacity
}

// function for getting node with the key provided
func (h *hashmap[K, V]) getNode(key K) (*node[K, V], error) {
	//1. generate hash
	hash := h.generateHash(key)

	//2. generate hash key
	hashKey := h.hashFn(hash)

	//3. search for the key
	curr := h.arrHead[hashKey]
	if curr == nil {
		return nil, errors.New("Node doesn't exist")
	}

	for curr != nil {
		if curr.key == key {
			return curr, nil
		}
		curr = curr.next
	}

	return nil, errors.New("Node doesn't exist")

}

// function for doubling the capacity of hashmap and reallocating all the key-value pairs.
func (h *hashmap[K, V]) rehash() {
	//1. Create new hash map with double capacity
	newhashmap := hashmap[K, V]{
		capacity:     2 * h.capacity,
		size:         h.size,
		arrHead:      make([]*node[K, V], 2*h.capacity),
		maxThreshold: h.maxThreshold,
	}

	//2. transfer all the key-value pairs
	for _, nodeHead := range h.arrHead {
		curr := nodeHead
		for curr != nil {
			newhashmap.set(curr.key, curr.value)
			curr = curr.next
		}
	}

	//3. udpate the hashmap with new one
	*h = newhashmap

}

// function for inserting value in hashmap
func (h *hashmap[K, V]) set(key K, value V) {

	//1. check if the node with given key exists or not
	targetNode, err := h.getNode(key)

	//2. If exists, change the value and return
	if err == nil {
		targetNode.value = value
		return
	}

	//3. check if rehashing is required
	currentThreshold := float32(h.size+1) / float32(h.capacity)
	if currentThreshold > h.maxThreshold {
		h.rehash()
	}

	//4. generate hash
	hash := h.generateHash(key)

	//5. generate hash key
	hashKey := h.hashFn(hash)

	//6. get the first node from the hash key position
	curr := h.arrHead[hashKey]

	//7. insert the new node in the beginning
	newNode := node[K, V]{
		key:   key,
		value: value,
		next:  curr,
	}
	h.arrHead[hashKey] = &newNode
	h.size++
}

// function for deleting value from hashmap
func (h *hashmap[K, V]) delete(key K) {
	//1. Check if key exists or not. If doesn't exist, return
	_, err := h.getNode(key)
	if err != nil {
		return
	}

	//2. generate hash
	hash := h.generateHash(key)
	//3. generate hash key
	hashKey := h.hashFn(hash)

	//4. get the head node
	headNode := h.arrHead[hashKey]

	//5. if key is present in the first node, update the head
	if headNode.key == key {
		h.arrHead[hashKey] = headNode.next
		h.size--
		return
	}

	//6. else find the node and remove it
	var prev *node[K, V]
	curr := headNode

	for curr.key != key {
		prev = curr
		curr = curr.next
	}

	prev.next = curr.next
	curr.next = nil
	h.size--
}

// function for getting value from hashmap
func (h *hashmap[K, V]) get(key K) (V, error) {
	//1. get the target node
	targetNode, err := h.getNode(key)

	//2. if targetNode not present, return error
	if err != nil {
		var defaultValue V
		return defaultValue, err
	}

	//3. return the value
	return targetNode.value, nil
}

// function for creating hashmap
func newHashMap[K comparable, V comparable]() *hashmap[K, V] {
	h := hashmap[K, V]{
		capacity:     10,
		size:         0,
		arrHead:      make([]*node[K, V], 10),
		maxThreshold: 0.7,
	}
	return &h
}

func main() {

	h := newHashMap[int, string]()

	h.set(1, "Mohit")
	h.set(2, "Uniyal")

	value, err := h.get(2)
	if err == nil {
		fmt.Println(value)
	}

}
