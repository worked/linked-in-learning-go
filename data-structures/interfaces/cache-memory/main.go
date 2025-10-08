// Write your answer here, and then test your code.
// Your job is to implement the Get and Set Cache methods.

package main

import (
	"errors"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type Key string
type Value string

type ValueWriter interface {
	Set(k Key, v Value)
}
type ValueReader interface {
	Get(k Key) (Value, error)
}

// Make any required changes to the Cache struct.
// Store the values in a map
type Cache struct {
	cacheMap map[Key]Value
}

// Create a new cache and initialize the cacheMap
func NewCache() *Cache {
	return &Cache{
		cacheMap: make(map[Key]Value),
	}
}

// Set the value in the cacheMap
func (c *Cache) Set(k Key, v Value) {
	c.cacheMap[k] = v
}

// Get the value from the cacheMap
func (c *Cache) Get(k Key) (Value, error) {
	v, ok := c.cacheMap[k]
	if !ok {
		return "", errors.New("value not found")
	}
	return v, nil
}

func WriteValues(w ValueWriter, keys []Key, values []Value) {
	for i, k := range keys {
		w.Set(k, values[i])
	}
	fmt.Println("cacheMap: ", w.(*Cache).cacheMap)
}

func ReadValues(r ValueReader, keys []Key) ([]Value, error) {
	values := make([]Value, len(keys))
	for i, k := range keys {
		v, err := r.Get(k)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}

	return values, nil
}

func main() {
	keys := []Key{"potato", "broccoli", "carrot", "banana", "potato", "banana"}
	values := []Value{"0.45", "1.5", "0.75", "2", "0.35", "2.4"}
	n := NewCache()
	WriteValues(n, keys, values)
	learnerResult, learnerError := ReadValues(n, keys)
	fmt.Println(learnerResult, learnerError)
	// expected result: [0.35 1.5 0.75 2.4 0.35 2.4] <nil>
}
