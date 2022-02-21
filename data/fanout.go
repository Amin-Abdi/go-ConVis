package data

import (
	"fmt"
	"sync"
)

func main() {
	status := make(chan string)

	go producer(status)

	fmt.Println(<-status)

}

func producer(a chan string) {
	var wg sync.WaitGroup
	channel := make(chan int, 4)

	wg.Add(4)

	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4

	go thread1(channel, &wg)
	go thread2(channel, &wg)
	go thread3(channel, &wg)
	go thread4(channel, &wg)

	wg.Wait()

	a <- "Done"

}

func thread1(a chan int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(<-a)
}

func thread2(a chan int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(<-a)
}

func thread3(a chan int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(<-a)
}

func thread4(a chan int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(<-a)
}
