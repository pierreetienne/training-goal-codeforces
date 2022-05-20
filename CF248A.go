package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF248A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF248A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF248A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF248A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF248A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF248A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF248A
Date: 18/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/248/A
Problem: CF248A
**/
func (in *CF248A) Run() {
	n := in.NextInt()
	l, r := 0, 0
	for i := 0; i < n; i++ {
		a, b := in.NextInt(), in.NextInt()
		l += a
		r += b
	}

	ll := int(math.Abs(float64(l) - float64(n)))
	rr := int(math.Abs(float64(r) - float64(n)))

	ans := 0
	if ll < l {
		ans += ll
	} else {
		ans += l
	}

	if rr < r {
		ans += rr
	} else {
		ans += r
	}
	fmt.Println(ans)
}

func NewCF248A(r *bufio.Reader) *CF248A {
	return &CF248A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF248A(bufio.NewReader(os.Stdin)).Run()
}
