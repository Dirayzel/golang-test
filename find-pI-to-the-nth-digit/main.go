package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Println("To how many decimal places you need to compute Pi?")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	if n >= 0 && n <= 100 {
		fmt.Println(n, math.Pi)
	} else {
		fmt.Println("Limits : 0 >= n <= 100")
	}
}
