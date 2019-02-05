package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func getPermission(p chan bool) {
	// Host which gives the permission to eat.
	p <- true
}

func (p Philo) eat(wg *sync.WaitGroup) {

	// Allowing philosphers to eat only 3 times
	// fmt.Println("starting to eat ", p.id+1)
	for i := 0; i < 3; i++ {

		fmt.Println("starting to eat ", p.id+1)
		p.leftCS.Lock()
		p.rightCS.Lock()
		
		time.Sleep(10 * time.Millisecond)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Println("finishing eating ", p.id+1)
	}

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Creating philosphers.
	philos := make([]*Philo, 5)

	// Randoming the order of picking up the chop sticks
	nums := []int{0, 1, 2, 3, 4}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	for k := 0; k < 5; k++ {
		j := nums[k]
		philos[j] = &Philo{j, CSticks[j], CSticks[(j+1)%5]}
	}

	for i := 0; i < 5; i += 2 {

		c1 := make(chan bool)
		c2 := make(chan bool)
		// Asking for the permission from the host.
		go getPermission(c1)
		go getPermission(c2)

		ph1 := <-c1
		ph2 := <-c2

		// Waiting for the permission to eat from the host
		if ph1 && ph2 {
			wg.Add(1)
			go philos[i].eat(&wg)

			if i+1 < 5 {
				wg.Add(1)
				go philos[i+1].eat(&wg)
			}

		}

		wg.Wait()
	}
}
