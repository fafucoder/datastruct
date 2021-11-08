/**
 * hash  表
 */

package main

import (
	"math"
	"sync"

	"github.com/OneOfOne/xxhash"
)

const expandFactor = 0.75

type HashTable struct {
	items        []*hashMap
	len          int
	capacity     int
	capacityMask int

	lock sync.Mutex
}

type hashMap struct {
	key   string
	value interface{}
	next  *hashMap
}

func NewHashTable(capacity int) *HashTable {
	defaultCapacity := 1 << 4
	if capacity < defaultCapacity {
		capacity = defaultCapacity
	} else {
		capacity = 1 << int(math.Ceil(math.Log2(float64(capacity))))
	}

	return &HashTable{
		items:        make([]*hashMap, capacity, capacity),
		capacity:     capacity,
		capacityMask: capacity - 1,
	}
}

func (c *HashTable) Add(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	index := c.hashIndex(key)
	element := c.items[index]
	if element == nil {
		c.items[index] = &hashMap{
			key:   key,
			value: value,
		}
	} else {
		var lastElement *hashMap
		for element != nil {
			if element.key == key {
				element.value = value
				return
			}

			lastElement = element
			element = element.next
		}

		lastElement.next = &hashMap{
			key:   key,
			value: value,
		}

		newLen := c.len + 1

		// 到达扩容界限
		if float64(newLen)/float64(c.capacity) >= expandFactor {
			newHashTable := new(HashTable)
			newHashTable.items = make([]*hashMap, 2*c.capacity, 2*c.capacity)
			newHashTable.capacity = 2 * c.capacity
			newHashTable.capacityMask = 2*c.capacity - 1
			for _, item := range c.items {
				for item != nil {
					newHashTable.Add(item.key, item.value)
					item = item.next
				}
			}

			c.items = newHashTable.items
			c.capacity = newHashTable.capacity
			c.capacityMask = newHashTable.capacityMask
		}

		c.len = newLen
	}
}

func (c *HashTable) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()

	index := c.hashIndex(key)
	element := c.items[index]

	for element != nil {
		if element.key == key {
			return element.value
		}

		element = element.next
	}

	return nil
}

func (c *HashTable) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	index := c.hashIndex(key)
	element := c.items[index]
	if element == nil {
		return
	}

	if element.key == key {
		c.items[index].next = element.next
		c.len = c.len - 1
		return
	}

	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			c.len = c.len - 1
			return
		}

		element = nextElement
		nextElement = nextElement.next
	}
}

func (c *HashTable) hashIndex(key string) int {
	hash := xxhash.Checksum64([]byte(key))

	return int(hash & uint64(c.capacityMask))
}
