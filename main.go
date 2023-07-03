package main

import (
	"bytes"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	buf := &bytes.Buffer{}
	ch := make(chan []byte, 50)
	go func() {
		for {
			buf.Write(<-ch)
			if buf.Len() < 1000000 {
				continue
			}
			io.Copy(os.Stdout, buf)
			buf.Reset()
		}
	}()
	for i := 1; i < math.MaxInt; i += 15 {
		ch <- s(i)
	}
}

func s(i int) []byte {
	return []byte(strconv.Itoa(i) + "\n" +
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
		"FizzBuzz\n")
}
