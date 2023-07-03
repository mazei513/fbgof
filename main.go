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

const base = 10

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
	sb.Grow(1000000)
	for i := int64(1); i < int64(*n); i += 15 {
		sb.WriteString(strconv.FormatInt(i, base))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+1, base))
		sb.WriteString("\nFizz\n")
		sb.WriteString(strconv.FormatInt(i+3, base))
		sb.WriteString("\nBuzz\nFizz\n")
		sb.WriteString(strconv.FormatInt(i+6, base))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+7, base))
		sb.WriteString("\nFizz\nBuzz\n")
		sb.WriteString(strconv.FormatInt(i+10, base))
		sb.WriteString("\nFizz\n")
		sb.WriteString(strconv.FormatInt(i+12, base))
		sb.WriteString("\n")
		sb.WriteString(strconv.FormatInt(i+13, base))
		sb.WriteString("\nFizzBuzz\n")
		if sb.Len() > 1000000 {
			os.Stdout.WriteString(sb.String())
			sb.Reset()
		}
	}
}
