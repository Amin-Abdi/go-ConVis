package data

func main() {
	consumer := make(chan int)

	worker1(consumer)
}

func worker1(a chan int) {
	a <- 99
}
