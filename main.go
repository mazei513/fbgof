package main

import (
	"bufio"
	"flag"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var n = flag.Int("n", math.MaxInt, "N")

const base = 10
const outBufMax = 5000

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
	sb := bufio.NewWriter(os.Stdout)
	for i := 1; i < *n; i += 15 {
		l, b := itoa(i)
		sb.Write(b)
		sb.Write([]byte("\n"))
		sb.Write(itoafast(l, i+1))
		sb.Write([]byte("\nFizz\n"))
		sb.Write(itoafast(l, i+3))
		sb.Write([]byte("\nBuzz\nFizz\n"))
		l, b = itoa(i + 6)
		sb.Write(b)
		sb.Write([]byte("\n"))
		sb.Write(itoafast(l, i+7))
		sb.Write([]byte("\nFizz\nBuzz\n"))
		l, b = itoa(i + 10)
		sb.Write(b)
		sb.Write([]byte("\nFizz\n"))
		sb.Write(itoafast(l, i+12))
		sb.Write([]byte("\n"))
		sb.Write(itoafast(l, i+13))
		sb.Write([]byte("\nFizzBuzz\n"))
	}
	sb.Flush()
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
const maxBuf = 64

var itoaBuf [maxBuf]byte

func itoa(u int) (int, []byte) {
	i := maxBuf
	us := uint(u)
	for us >= 100 {
		is := us % 100 * 2
		us /= 100
		i -= 2
		itoaBuf[i+1] = smallsString[is+1]
		itoaBuf[i+0] = smallsString[is+0]
	}

	is := us * 2
	i--
	itoaBuf[i] = smallsString[is+1]
	if us >= 10 {
		i--
		itoaBuf[i] = smallsString[is]
	}
	return i, itoaBuf[i:]
}

func itoafast(l, u int) []byte {
	i := maxBuf
	us := uint(u)
	is := us % 100 * 2
	us /= 100
	i -= 2
	itoaBuf[i+1] = smallsString[is+1]
	itoaBuf[i+0] = smallsString[is+0]

	return itoaBuf[l:]
}
