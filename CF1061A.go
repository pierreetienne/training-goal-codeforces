package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1061A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1061A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1061A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1061A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1061A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1061A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1061A
Date: 17/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1061/A
Problem: CF1061A
**/
func (in *CF1061A) Run() {
	n, S := in.NextInt(), in.NextInt()
	coins := int(math.Floor(float64(S) / float64(n)))
	if coins*n == S {
		fmt.Println(coins)
	} else {
		fmt.Println(coins + 1)
	}

}

func NewCF1061A(r *bufio.Reader) *CF1061A {
	return &CF1061A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1061A(bufio.NewReader(os.Stdin)).Run()
}
