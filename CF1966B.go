package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1966B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1966B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1966B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1966B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1966B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1966B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1966B
Date: 17/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1696/B
Problem: CF1966B
**/
func (in *CF1966B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		zero := 0
		last := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if val != 0 && last == 0 {
				zero++
			}
			last = val
		}
		if zero > 2 {
			zero = 2
		}
		fmt.Println(zero)
	}
}

func NewCF1966B(r *bufio.Reader) *CF1966B {
	return &CF1966B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1966B(bufio.NewReader(os.Stdin)).Run()
}
