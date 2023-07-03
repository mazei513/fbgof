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
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("\n")
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.Itoa(i + 3))
		sb.WriteString("\n")
		sb.WriteString("Buzz\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.Itoa(i + 6))
		sb.WriteString("\n")
		sb.WriteString(strconv.Itoa(i + 7))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString("Buzz\n")
		sb.WriteString(strconv.Itoa(i + 10))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.Itoa(i + 12))
		sb.WriteString("\n")
		sb.WriteString(strconv.Itoa(i + 13))
		sb.WriteString("\n")
		sb.WriteString("FizzBuzz\n")
		if sb.Len() > 500000 {
			os.Stdout.WriteString(sb.String())
			sb.Reset()
		}
	}
}
