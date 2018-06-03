package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MagicMinute is our own minute to make execution faster.
const MagicMinute time.Duration = time.Millisecond

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	users := make(chan struct{}, 8)

	for i := 1; i < 26; i++ {
		wg.Add(1)
		go User(i, users, &wg)
	}

	wg.Wait()
	fmt.Println("The place is empty, let's close up and go to the beach!")
}

// User is user action goroutine.
func User(index int, users chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		// remove user from pool
		<-users
		wg.Done()
	}()

	select {
	case users <- struct{}{}:
	default:
		fmt.Printf("Tourist %d waiting for turn\n", index)
		users <- struct{}{}
	}
	fmt.Printf("Tourist %d is online\n", index)
	delay := random(15, 120)
	time.Sleep(time.Duration(delay) * MagicMinute)
	fmt.Printf("Tourist %d is done, having spent %d minutes online.\n", index, delay)
}

// Random returns random number in [min, max).
func random(min, max int) int {
	return min + rand.Intn(max-min)
}
