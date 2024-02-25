package main

import (
	"fmt"
	"sync"
)

const numPhilosophers = 5
const numMealsPerPhilosopher = 3
const maxConcurrentEaters = 2

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
	mealsEaten     int
}

type Host struct {
	permission chan bool
}

func NewHost() *Host {
	return &Host{permission: make(chan bool, maxConcurrentEaters)}
}

func (h *Host) getPermission() {
	h.permission <- true
}

func (h *Host) releasePermission() {
	<-h.permission
}

func (p *Philosopher) eat() {
	p.host.getPermission()
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()

	fmt.Printf("starting to eat %d\n", p.id)
	p.mealsEaten++
	fmt.Printf("finishing eating %d\n", p.id)

	p.rightChopstick.Unlock()
	p.leftChopstick.Unlock()
	p.host.releasePermission()
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMealsPerPhilosopher; i++ {
		p.eat()
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	host := NewHost()

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := range philosophers {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           host,
		}
	}

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go p.dine(&wg)
	}
	wg.Wait()
}
