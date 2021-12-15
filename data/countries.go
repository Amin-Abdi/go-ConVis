package data

import (
	"fmt"
	"math/rand"
)

func GetCountries() {
	myInt := rand.Int()

	myCountries := []string{"Canada", "USA", "Australia", "England"}

	fmt.Println(myInt)
	fmt.Println(myCountries)
}
