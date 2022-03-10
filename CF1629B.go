package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1629B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1629B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1629B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1629B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1629B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1629B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1629B
Date: 8/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1629/B
Problem: CF1629B
**/
func (in *CF1629B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		l, r, k := in.NextInt(), in.NextInt(), in.NextInt()

		total := r - (l - 1)
		even := int(math.Ceil((float64(r) - (float64(l) - 1.)) / 2.))
		x := l%2 != 0
		y := r%2 != 0
		if x && y {
			even--
		}

		odd := total - even

		if (even > 0 && odd <= k) || (total == 1 && r != 1 && l != 1) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}

func NewCF1629B(r *bufio.Reader) *CF1629B {
	return &CF1629B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1629B(bufio.NewReader(os.Stdin)).Run()
}
