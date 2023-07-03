package main

import (
	"bytes"
	"math"
	"os"
	"strconv"
)

func main() {
	buf := bytes.Buffer{}
	ch := make(chan string, 50)
	ch2 := make(chan []byte, 2)
	go func() {
		for {
			buf.WriteString(<-ch)
			if buf.Len() < 100000 {
				continue
			}
			ch2 <- buf.Bytes()
			buf.Truncate(0)
		}
	}()
	go func() {
		for {
			o(<-ch2)
		}
	}()
	for i := 1; i < math.MaxInt; i += 15 {
		ch <- s(i)
	}
}

func s(i int) string {
	return strconv.Itoa(i) + "\n" +
		strconv.Itoa(i+1) + "\n" +
		"Fizz\n" +
		strconv.Itoa(i+3) + "\n" +
		"Buzz\n" +
		"Fizz\n" +
		strconv.Itoa(i+6) + "\n" +
		strconv.Itoa(i+7) + "\n" +
		"Fizz\n" +
		"Buzz\n" +
		strconv.Itoa(i+10) + "\n" +
		"Fizz\n" +
		strconv.Itoa(i+12) + "\n" +
		strconv.Itoa(i+13) + "\n" +
		"FizzBuzz\n"
	// return fmt.Sprintf("%d\n%d\nFizz\n%d\nBuzz\nFizz\n%d\n%d\nFizz\nBuzz\n%d\nFizz\n%d\n%d\nFizzBuzz\n",
	// 	i, i+1, i+3, i+6, i+7, i+10, i+12, i+13)
}
func o(b []byte) {
	os.Stdout.Write(b)
}
