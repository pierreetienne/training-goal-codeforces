package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1191A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1191A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1191A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1191A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1191A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1191A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1191A
Date: 30/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1191/A
Problem: CF1191A
**/
func (in *CF1191A) Run() {
	n := in.NextInt()
	current := F(n)
	next := F(n + 1)
	next2 := F(n + 2)

	if next < next2 && next < current {
		fmt.Println("1", string(next))
	} else if next2 < next && next2 < current {
		fmt.Println("2", string(next2))
	} else {
		fmt.Println("0", string(current))
	}
}

func F(n int) uint8 {
	if n%4 == 1 {
		return 'A'
	} else if n%4 == 3 {
		return 'B'
	} else if n%4 == 2 {
		return 'C'
	} else {
		return 'D'
	}
}

func NewCF1191A(r *bufio.Reader) *CF1191A {
	return &CF1191A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1191A(bufio.NewReader(os.Stdin)).Run()
}
