package main

import (
	"fmt"
	"sync"
)

const VISITORS_NUMBER = 10

type Person struct {
	status string
	id     int
}

type Turn struct {
	turn  chan *Person
	mutex sync.Mutex
}

func (t *Turn) add(visitor *Person) {
	defer t.mutex.Unlock()
	t.mutex.Lock()
	visitor.status = "in turn"
	t.turn <- visitor
}

func barber() {

}

func client(turn Turn, client_id int, wg *sync.WaitGroup) {
	defer wg.Done()
	visitor := Person{"wait", client_id}
	go turn.add(&visitor)
	select {
	case <-t.turn:

	}
	fmt.Printf("%s\n", visitor.status)
}

func main() {
	var wg sync.WaitGroup

	turn := Turn{}
	turn.turn = make(chan *Person)

	for i := 1; i <= VISITORS_NUMBER; i++ {
		wg.Add(1)
		go client(turn, 1, &wg)
	}

	wg.Wait()
	fmt.Println("end")
}
