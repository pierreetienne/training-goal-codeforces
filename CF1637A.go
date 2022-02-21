package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1637A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1637A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1637A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1637A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1637A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1637A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1637A
Date: 18/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1637/A
Problem: CF1637A
**/
func (in *CF1637A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int64, n)
		ans := true
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt64()
		}

		for i:=1;i<n;i++{
			if arr[i-1]>arr[i]{
				ans = false
			}
		}
		if !ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1637A(r *bufio.Reader) *CF1637A {
	return &CF1637A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1637A(bufio.NewReader(os.Stdin)).Run()
}

/*
1
4
999999999 999999999 999999998 1000000000

*/
