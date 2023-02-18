package data

import "math"

func main1() {
	cha1 := make(chan string)
	myVal := math.Log10(1148)

	go handler(cha1, myVal)

	<-cha1
}

func handler(c chan string, a float64) {
	converted := int(a)

	if converted%2 == 0 {
		c <- "EVEN"
	} else {
		c <- "ODD"
	}

}
