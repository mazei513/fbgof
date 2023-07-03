package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sb := &strings.Builder{}
	sb.Grow(1000000)
	for i := 1; i < math.MaxInt; i += 15 {
		sb.WriteString(strconv.Itoa(i) + "\n" +
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
		if sb.Len() > 500000 {
			os.Stdout.WriteString(sb.String())
			sb.Reset()
		}
	}
}
