package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < math.MaxInt; i++ {
		str := ""
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = fmt.Sprint(i)
		}
		fmt.Println(str)
	}
}