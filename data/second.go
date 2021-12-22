package data

func main() {
	first := make(chan int)
	third := make(chan int)

	go sender(first, "Hasicas", third)

	<-first
}

func sender(c chan int, k string, a chan int) {
	c <- 99
}
