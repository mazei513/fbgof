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
	interBuf := make([]byte, 0, 19*15)
	for i := 1; i < *n; i += 15 {
		interBuf = append(interBuf, itoa(i)...)
		interBuf = append(interBuf, '\n')
		interBuf = append(interBuf, itoa(i+1)...)
		interBuf = append(interBuf, []byte("\nFizz\n")...)
		interBuf = append(interBuf, itoa(i+3)...)
		interBuf = append(interBuf, []byte("\nBuzz\nFizz\n")...)
		interBuf = append(interBuf, itoa(i+6)...)
		interBuf = append(interBuf, '\n')
		interBuf = append(interBuf, itoa(i+7)...)
		interBuf = append(interBuf, []byte("\nFizz\nBuzz\n")...)
		interBuf = append(interBuf, itoa(i+10)...)
		interBuf = append(interBuf, []byte("\nFizz\n")...)
		interBuf = append(interBuf, itoa(i+12)...)
		interBuf = append(interBuf, '\n')
		interBuf = append(interBuf, itoa(i+13)...)
		interBuf = append(interBuf, []byte("\nFizzBuzz\n")...)
		_, err := sb.Write(interBuf)
		if err != nil {
			panic(err)
		}
		interBuf = interBuf[:0]
	}
	err := sb.Flush()
	if err != nil {
		panic(err)
	}
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
const maxBuf = 19

var itoaBuf [maxBuf]byte
var lastL int
var lastUs uint

func itoa(u int) []byte {
	i := maxBuf
	us := uint(u)
	if us >= 100 {
		is := us % 100 * 2
		us /= 100
		i -= 2
		itoaBuf[i+1] = smallsString[is+1]
		itoaBuf[i+0] = smallsString[is+0]
		if lastUs == us {
			return itoaBuf[lastL:]
		}
		lastUs = us
	}
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
	lastL = i
	return itoaBuf[i:]
}
