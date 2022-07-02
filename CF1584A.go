package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1584A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1584A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1584A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1584A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1584A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1584A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1584A
Date: 22/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1584/A
Problem: CF1584A
**/
func (in *CF1584A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := float64(in.NextInt()),float64( in.NextInt())
		if a-b == 0 {
			fmt.Println(1, -1)
		} else {
			n := float64(-10)
			for {
				y := ((n * (a * b)) - (a * b) - (b * b)) / (a - b)
				x := ((a * b) - (a * y)) / b

				p := float64(int(x))
				q := float64(int(y))

				r1 := (p + q) / (a + b)
				r2 := ((b * p) + (a * q)) / (a * b)
				fmt.Println("-:>" , r1, r2)
				if math.Abs(r1 - r2) < 1e-6 {
					fmt.Println(x, y)
					break
				} else {
					n++
					if n == 0 {
						n =1
					}
					if n > 10 {
						break
					}
					fmt.Println("pailas")
				}
			}
		}

	}
}

func NewCF1584A(r *bufio.Reader) *CF1584A {
	return &CF1584A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1584A(bufio.NewReader(os.Stdin)).Run()
}
