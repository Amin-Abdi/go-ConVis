package data

func main4() {
	first := make(chan int)
	second := make(chan string)

	go go1(first)
	go go2(first, second)

	<-second

}

func go1(c chan int) {
	c <- 99
}

func go2(a chan int, b chan string) {
	<-a
	b <- "variable 2"
}
