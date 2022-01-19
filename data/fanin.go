package data

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	consumer := make(chan int, 1)
	wg.Add(3)

	go worker1(consumer, &wg)
	go worker2(consumer, &wg)
	go worker3(consumer, &wg)
	wg.Wait()

	<-consumer
	<-consumer
	<-consumer

}
