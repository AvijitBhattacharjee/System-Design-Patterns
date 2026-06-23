package main

import (
	"fmt"
	"sync"
	"time"
)


type RateLimiter struct {
	requestCount, capacity int
	request map[int64]int
	windowSize int64
	mu sync.Mutex
}

func (r *RateLimiter)Allow() bool {
	
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now().Unix()
	for timeStamp := range r.request {
		if now-timeStamp >= r.windowSize {
			delete(r.request, timeStamp)
		}
	}

	r.requestCount = 0

	for timeStamp, count := range r.request {
		if now-timeStamp < r.windowSize {
			r.requestCount+= count
		}
	}

	if r.capacity < r.requestCount {return false}

	r.request[now]++
	return true
}

func main() {
	fmt.Println("Implementing rate limiter")

	rl := &RateLimiter{
		capacity: 5,
		windowSize: 5,
		request: make(map[int64]int),
	}

	for i:=0;i<=10;i++ {

		fmt.Println("Request no allow ->\n", i, rl.Allow())
		time.Sleep(200*time.Millisecond)
	}

	fmt.Println("\nWaiting for window expiry...")

	time.Sleep(6 * time.Second)

	for i := 11; i <= 15; i++ {

		fmt.Printf("Request %d -> %v\n", i, rl.Allow())

		time.Sleep(500 * time.Millisecond)
	}

	

	
}