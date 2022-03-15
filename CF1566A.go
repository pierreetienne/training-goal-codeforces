package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1566A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1566A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1566A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1566A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1566A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1566A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1566A
Date: 13/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1566/A
Problem: CF1566A
**/
func (in *CF1566A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, s := in.NextInt(), in.NextInt()
		pos := int(math.Ceil(float64(n)/2.0)) - 1
		if n-pos != 0 {
			nums := int(math.Floor(float64(s / (n - pos))))
			fmt.Println(nums)
		} else {
			fmt.Println(s)
		}
	}
}

func NewCF1566A(r *bufio.Reader) *CF1566A {
	return &CF1566A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1566A(bufio.NewReader(os.Stdin)).Run()
}
