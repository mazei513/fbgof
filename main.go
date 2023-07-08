package main

import (
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
	const interBufL = 2048 * 10
	const interL = 19 * 15
	interBuf := make([]byte, interBufL)
	fb1 := []byte("Fizz\n")
	fb2 := []byte("Buzz\nFizz\n")
	fb3 := []byte("Fizz\nBuzz\n")
	fb4 := []byte("FizzBuzz\n")
	var bn int
	for i := 1; i < *n; i += 15 {
		bn = lenCpy(bn, interBuf, itoa(i))
		bn = lenCpy(bn, interBuf, itoa(i+1))
		bn = lenCpy(bn, interBuf, fb1)
		bn = lenCpy(bn, interBuf, itoa(i+3))
		bn = lenCpy(bn, interBuf, fb2)
		bn = lenCpy(bn, interBuf, itoa(i+6))
		bn = lenCpy(bn, interBuf, itoa(i+7))
		bn = lenCpy(bn, interBuf, fb3)
		bn = lenCpy(bn, interBuf, itoa(i+10))
		bn = lenCpy(bn, interBuf, fb1)
		bn = lenCpy(bn, interBuf, itoa(i+12))
		bn = lenCpy(bn, interBuf, itoa(i+13))
		bn = lenCpy(bn, interBuf, fb4)
		if bn > interBufL-interL {
			_, err := os.Stdout.Write(interBuf[:bn])
			if err != nil {
				panic(err)
			}
			bn = 0
		}
	}
	_, err := os.Stdout.Write(interBuf[:bn])
	if err != nil {
		panic(err)
	}
}

func lenCpy(ln int, dst, src []byte) int {
	copy(dst[ln:], src)
	return ln + len(src)
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
	itoaBuf[i-1] = '\n'
	i--
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
