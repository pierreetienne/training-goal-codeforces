package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1675A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1675A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1675A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1675A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1675A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1675A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1675A
Date: 9/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1675/A
Problem: CF1675A
**/
func (in *CF1675A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b, c, x, y := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()

		if a >= x {
			a -= x
			a = 0
			x = 0
		} else {
			x-= a
			a = 0
			if x <= c {
				c -= x
				x = 0
			}
		}

		if b >= y {
			a -= y
			b = 0
			y = 0
		} else {
			y-= b
			b = 0
			if y <= c {
				c -= y
				y = 0
			}

		}

		if x == 0 && y == 0{
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}

	}

}

func NewCF1675A(r *bufio.Reader) *CF1675A {
	return &CF1675A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1675A(bufio.NewReader(os.Stdin)).Run()
}
