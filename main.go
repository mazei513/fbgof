package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var n = flag.Int("n", 1<<60, "N")

const loopsPerWorker = 160000
const linesPerLoop = 15
const bufSize = loopsPerWorker * linesPerLoop * maxBuf

type job struct{ start, end int }

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
	nWorkers := runtime.NumCPU() - 1
	jobs := make([]chan job, nWorkers)
	outs := make([]chan []byte, nWorkers)
	done := make(chan struct{})
	for i := 0; i < nWorkers; i++ {
		jobs[i] = make(chan job)
		outs[i] = make(chan []byte)
		go fb(jobs[i], outs[i])
	}
	go out(outs, done)
	cnt := 1
	rotate := 0
	for cnt < *n {
		end := cnt + loopsPerWorker*linesPerLoop
		jobs[rotate] <- struct{ start, end int }{cnt, end}
		rotate++
		rotate %= nWorkers
		cnt = end
	}
	for i := 0; i < nWorkers; i++ {
		close(jobs[(i+rotate)%nWorkers])
	}
	<-done
}

func out(jobs []chan []byte, done chan struct{}) {
	for {
		for _, j := range jobs {
			o, more := <-j
			if !more {
				done <- struct{}{}
			}
			os.Stdout.Write(o)
		}
	}
}

func fb(in chan job, out chan []byte) {
	fb1 := []byte("Fizz\n")
	fb2 := []byte("Buzz\nFizz\n")
	fb3 := []byte("Fizz\nBuzz\n")
	fb4 := []byte("FizzBuzz\n")
	interBuf := make([]byte, 0, bufSize)
	ib := &itoaBuf{}

	for {
		j, more := <-in
		if !more {
			break
		}
		for i := j.start; i < j.end; i += 15 {
			interBuf = append(interBuf, ib.itoa(i)...)
			interBuf = append(interBuf, ib.itoa(i+1)...)
			interBuf = append(interBuf, fb1...)
			interBuf = append(interBuf, ib.itoa(i+3)...)
			interBuf = append(interBuf, fb2...)
			interBuf = append(interBuf, ib.itoa(i+6)...)
			interBuf = append(interBuf, ib.itoa(i+7)...)
			interBuf = append(interBuf, fb3...)
			interBuf = append(interBuf, ib.itoa(i+10)...)
			interBuf = append(interBuf, fb1...)
			interBuf = append(interBuf, ib.itoa(i+12)...)
			interBuf = append(interBuf, ib.itoa(i+13)...)
			interBuf = append(interBuf, fb4...)
		}
		out <- interBuf
		interBuf = interBuf[:0]
	}
	close(out)
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

type itoaBuf struct {
	buf    [maxBuf]byte
	lastL  int
	lastUs uint
}

func (ib *itoaBuf) itoa(u int) []byte {
	i := maxBuf
	ib.buf[i-1] = '\n'
	i--
	us := uint(u)
	if us >= 100 {
		is := us % 100 * 2
		us /= 100
		i -= 2
		ib.buf[i+1] = smallsString[is+1]
		ib.buf[i+0] = smallsString[is+0]
		if ib.lastUs == us {
			return ib.buf[ib.lastL:]
		}
		ib.lastUs = us
	}
	for us >= 100 {
		is := us % 100 * 2
		us /= 100
		i -= 2
		ib.buf[i+1] = smallsString[is+1]
		ib.buf[i+0] = smallsString[is+0]
	}

	is := us * 2
	i--
	ib.buf[i] = smallsString[is+1]
	if us >= 10 {
		i--
		ib.buf[i] = smallsString[is]
	}
	ib.lastL = i
	return ib.buf[i:]

}
