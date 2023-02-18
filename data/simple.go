package data

func main7() {
	first := make(chan int)

	go sender(first)

	<-first

}

func sender(c chan int) {
	c <- 99
}
