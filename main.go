package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	buf := bytes.Buffer{}
	ch := make(chan string, 50)
	go func() {
		for {
			buf.WriteString(<-ch)
			if buf.Len() < 1000 {
				continue
			}
			os.Stdout.Write(buf.Bytes())
			buf.Truncate(0)
		}
	}()
	for i := 1; i < math.MaxInt; i += 15 {
		str := fmt.Sprintf("%d\n%d\nFizz\n%d\nBuzz\nFizz\n%d\n%d\nFizz\nBuzz\n%d\nFizz\n%d\n%d\nFizzBuzz\n",
			i, i+1, i+3, i+6, i+7, i+10, i+12, i+13)
		ch <- str
	}
}
