package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1162A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1162A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1162A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1162A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1162A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1162A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1162A
Date: 9/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1162/A
Problem: CF1162A
**/
func (in *CF1162A) Run() {
	n, h, m := in.NextInt(), in.NextInt(), in.NextInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = 10000000007
	}
	for i := 0; i < m; i++ {
		l, r, m := in.NextInt(), in.NextInt(), in.NextInt()
		for ; l <= r; l++ {
			arr[l-1] = int(math.Min(float64(h), math.Min(float64(arr[l-1]), float64(m))))
		}
	}

	sum := int64(0)
	for i := 0; i < n; i++ {
		if arr[i] == 10000000007 {
			arr[i] = h
		}
		sum += int64(arr[i]) * int64(arr[i])
	}

	fmt.Println(sum)
}

func NewCF1162A(r *bufio.Reader) *CF1162A {
	return &CF1162A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1162A(bufio.NewReader(os.Stdin)).Run()
}
