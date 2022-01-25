package data

import (
	"fmt"
	"sync"
)

//Use only this class for testing use cases don't change the other test cases.
func main() {
	var wg sync.WaitGroup

	cons := make(chan int, 1)
	prod := make(chan int, 1)

	wg.Add(1)
	go mine(cons, &wg)
	wg.Wait()

	wg.Add(1)
	go mine2(cons, &wg, prod)
	wg.Wait()

	fmt.Println(<-prod)
}

func mine(a chan int, w *sync.WaitGroup) {
	a <- 10
	w.Done()
}

func mine2(in chan int, w *sync.WaitGroup, out chan int) {

	val := <-in
	out <- val
	close(out)

	w.Done()
}
