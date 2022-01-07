package data

func main() {
	first := make(chan int)

	go sender(first)

	<-first

}

func sender(c chan int) {
	c <- 99
}
