package data

func main() {
	table := make(chan int)

	go player1(table)
	go player2(table)

	<-table
}

func player1(c chan int) {
	c <- 1
	<-c
	c <- 1
}

func player2(b chan int) {
	<-b
	b <- 2
	<-b
	b <- 2
}
