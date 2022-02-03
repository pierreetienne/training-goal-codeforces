package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF900A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF900A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF900A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF900A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF900A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF900A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF900A
Date: 1/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/900/A
Problem: CF900A
**/
func (in *CF900A) Run() {
	countNeg := 0
	countPos := 0
	for n := in.NextInt(); n > 0; n-- {
		x, _ := in.NextInt64(), in.NextInt64()
		if x > 0 {
			countPos++
		} else {
			countNeg++
		}
	}

	if countNeg == 1 || countPos == 1 || (countNeg > 0 && countPos == 0) || (countNeg == 0 && countPos > 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func NewCF900A(r *bufio.Reader) *CF900A {
	return &CF900A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF900A(bufio.NewReader(os.Stdin)).Run()
}
