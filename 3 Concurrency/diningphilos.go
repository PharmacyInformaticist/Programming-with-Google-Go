package main

import (
	"fmt"
	"sync"
)

// Chopsticks and Philosophers
type ChopS struct{ sync.Mutex }

// There should be 5 philosophers sharing chopsticks,
// with one chopstick between each adjacent pair of philosophers.
type Philo struct {
	leftCS, rightCS *ChopS
}

// Philosopher Eat Method
// Each philosopher should eat only 3 times
// (not in an infinite loop as we did in lecture)
func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

// Initialization in Main
func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i],
			CSticks[(i+1)%5]}
	}

	// Start the Dining in Main
	// needs wait groups
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
	// Deadlock Problem
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("eating")
	p.rightCS.Unlock()
	p.leftCS.Unlock()
	// Deadlock Solution
	philos[i] = &Philo{CSticks[i],
		CSticks[(i+1)%5]}
}
