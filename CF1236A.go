package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1236A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1236A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1236A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1236A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1236A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1236A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots
Run solve the problem CF1236A
Date: 22/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1236/A
Problem: CF1236A
**/
func (in *CF1236A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		i, j, p := in.NextInt(), in.NextInt(), in.NextInt()
		ans := 0

		for a, b, c := i, j, p; ; {
			x := false
			if a > 0 && b >= 2 {
				ans += 3
				a--
				b -= 2
				x = true
			} else if b > 0 && c >= 2 {
				ans += 3
				b--
				c -= 2
				x = true
			}
			if !x {
				break
			}
		}

		beta := 0

		for a, b, c := i, j, p; ; {
			x := false
			if b > 0 && c >= 2 {
				beta += 3
				b--
				c -= 2
				x = true
			} else if a > 0 && b >= 2 {
				beta += 3
				a--
				b -= 2
				x = true
			}
			if !x {
				break
			}
		}

		fmt.Println(int(math.Max(float64(ans), float64(beta))))
	}
}

func NewCF1236A(r *bufio.Reader) *CF1236A {
	return &CF1236A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1236A(bufio.NewReader(os.Stdin)).Run()
}
