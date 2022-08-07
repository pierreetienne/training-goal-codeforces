package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1716A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1716A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1716A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1716A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1716A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1716A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1716A
Date: 5/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1716/A
Problem: CF1716A
**/
func (in *CF1716A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		ans := int64(math.Ceil(float64(n) / 3.))
		if n == 1 {
			fmt.Println(2)
		} else {
			fmt.Println(ans)
		}
	}
}

func NewCF1716A(r *bufio.Reader) *CF1716A {
	return &CF1716A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1716A(bufio.NewReader(os.Stdin)).Run()
}
