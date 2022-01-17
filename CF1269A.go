package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1269A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1269A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1269A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1269A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1269A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1269A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1269A
Date: 15/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1269/A
Problem: CF1269A
**/
func (in *CF1269A) Run() {
	n := in.NextInt()

	ws := false
	a := n
	b := 0
	for !ws {
		if composite(b) && composite(a) {
			fmt.Println(a, b)
			ws = true
		} else {
			a++
			b++
		}
	}

}

func composite(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func NewCF1269A(r *bufio.Reader) *CF1269A {
	return &CF1269A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1269A(bufio.NewReader(os.Stdin)).Run()
}
