package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MagicSecond is our own second to make execution faster.
const MagicSecond time.Duration = time.Millisecond

var (
	dishes  = [...]string{"chorizo", "chopitos", "pimientos de padrón", "croquetas", "patatas bravas"}
	persons = [...]string{"Alice", "Bob", "Charlie", "Dave"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var dishesWg, personsWg sync.WaitGroup

	fmt.Println("Bon appétit!")
	morsels := make(chan string)
	for _, dish := range dishes {
		dishesWg.Add(1)
		go DishProducer(dish, morsels, &dishesWg)
	}

	for _, person := range persons {
		personsWg.Add(1)
		go PersonConsumer(person, morsels, &personsWg)
	}

	dishesWg.Wait()
	close(morsels)

	personsWg.Wait()
	fmt.Println("That was delicious!")
}

// DishProducer is a goroutine which populates morsels chanel.
func DishProducer(dish string, morsels chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	morselsCount := random(5, 10)
	for i := 0; i < morselsCount; i++ {
		morsels <- dish
	}
}

// PersonConsumer is a goroutine to perform person's actions.
func PersonConsumer(person string, morsels chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		dish, ok := <-morsels
		if !ok {
			return
		}
		fmt.Printf("%s is enjoying some %s\n", person, dish)
		time.Sleep(time.Duration(random(30, 180)) * MagicSecond)
	}
}

// Random returns random number in [min, max).
func random(min, max int) int {
	return min + rand.Intn(max-min)
}
