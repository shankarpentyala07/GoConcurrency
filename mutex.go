package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// Create an instance of SafeCounter with an empty map.
	c := SafeCounter{v: make(map[string]int)}

	// Spawn 1000 goroutines to concurrently increment the counter for the key "somekey".
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// Allow some time for the goroutines to complete before printing the result.
	time.Sleep(time.Second)

	// Print the final value of the counter for the key "somekey".
	fmt.Println(c.Value("somekey"))
}


/*
Explanation:

SafeCounter Struct:

SafeCounter is a struct that represents a counter with a mutex to make it safe for concurrent access.
mu is a sync.Mutex used to synchronize access to the shared data.
v is a map where the keys are strings and the values are integers. It represents the counter.
Inc Method:

Inc is a method of SafeCounter that increments the counter for the given key.
It locks the mutex (c.mu) before modifying the shared map (c.v) to ensure exclusive access.
It then unlocks the mutex to allow other goroutines to access the map.
Value Method:

Value is a method of SafeCounter that returns the current value of the counter for the given key.
Similar to Inc, it locks the mutex before accessing the shared map and defers the unlock until the function returns.
Main Function:

In the main function, an instance of SafeCounter is created with an empty map.
1000 goroutines are spawned, each calling the Inc method to increment the counter for the key "somekey" concurrently.
time.Sleep(time.Second) is used to give the goroutines some time to complete before proceeding.
The final value of the counter for the key "somekey" is printed, and it should be 1000 if there were no race conditions. The sync.Mutex ensures that the increment operation is atomic and safe for concurrent use.
*/
