package main

import (
	"flag"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var n = flag.Int("n", math.MaxInt, "N")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	sb := &strings.Builder{}
	for i := 1; i < *n; i += 15 {
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
