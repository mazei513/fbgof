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
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

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
	for i := int64(1); i < int64(*n); i += 15 {
		sb.WriteString(strconv.FormatInt(i, 16))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+1, 16))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.FormatInt(i+3, 16))
		sb.WriteString("\n")
		sb.WriteString("Buzz\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.FormatInt(i+6, 16))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+7, 16))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString("Buzz\n")
		sb.WriteString(strconv.FormatInt(i+10, 16))
		sb.WriteString("\n")
		sb.WriteString("Fizz\n")
		sb.WriteString(strconv.FormatInt(i+12, 16))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+13, 16))
		sb.WriteString("\n")
		sb.WriteString("FizzBuzz\n")
		if sb.Len() > 1000000 {
			os.Stdout.WriteString(sb.String())
			sb.Reset()
		}
	}
}
