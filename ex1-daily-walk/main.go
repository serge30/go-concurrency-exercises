package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MagicSecond is our own second to make execution faster.
const MagicSecond time.Duration = time.Millisecond

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg, alarmWg sync.WaitGroup

	fmt.Println("Let's go for a walk!")
	wg.Add(2)
	go Prepare("Bob", &wg)
	go Prepare("Alice", &wg)
	wg.Wait()

	fmt.Println("Arming alarm.")
	alarmWg.Add(1)
	go Alarm(&alarmWg)

	wg.Add(2)
	go PutShoes("Bob", &wg)
	go PutShoes("Alice", &wg)
	wg.Wait()

	fmt.Println("Exiting and locking the door.")
	alarmWg.Wait()
}

// Prepare is a function for a preparation.
func Prepare(person string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s started getting ready\n", person)
	delay := random(60, 90)
	time.Sleep(time.Duration(delay) * MagicSecond)
	fmt.Printf("%s spent %d seconds getting ready\n", person, delay)
}

// PutShoes is a function for a putting shoes.
func PutShoes(person string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s started putting on shoes\n", person)
	delay := random(35, 45)
	time.Sleep(time.Duration(delay) * MagicSecond)
	fmt.Printf("%s spent %d seconds putting on shoes\n", person, delay)
}

// Alarm is alram countdownd function.
func Alarm(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Alarm is counting down.")
	time.Sleep(60 * MagicSecond)
	fmt.Println("Alarm is armed.")
}

// Random returns random number in [min, max).
func random(min, max int) int {
	return min + rand.Intn(max-min)
}
