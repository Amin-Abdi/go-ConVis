package data

import (
	"fmt"
	"sync"
)

func main2() {
	var wg sync.WaitGroup
	consumer := make(chan int, 3)
	result := make(chan string, 1)
	wg.Add(3)

	go worker1(consumer, &wg)
	go worker2(consumer, &wg)
	go worker3(consumer, &wg)

	wg.Wait()
	close(consumer)

	wg.Add(1)
	go getResult(consumer, &wg, result)
	wg.Wait()

	fmt.Println("Sum:", <-result)
	fmt.Println("Main Finished")

}

func worker1(a chan int, w *sync.WaitGroup) {
	defer w.Done()
	a <- 10
}

func worker2(b chan int, w *sync.WaitGroup) {
	b <- 20
	defer w.Done()
}
func worker3(c chan int, w *sync.WaitGroup) {
	c <- 90
	defer w.Done()
}

func getResult(in chan int, w *sync.WaitGroup, out chan string) {
	defer w.Done()
	a := <-in
	b := <-in
	c := <-in

	mySum := a + b + c

	ans := fmt.Sprintf("The sum is %v\n", mySum)

	out <- ans
	close(out)
}
