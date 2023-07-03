package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	ch := make(chan string, 50)
	go func() {
		for {
			out(<-ch)
		}
	}()
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
		ch <- str
	}
}

func out(str string) {
	os.Stdout.WriteString(str + "\n")
}
