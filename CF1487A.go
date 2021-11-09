package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1487A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1487A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1487A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1487A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1487A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1487A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1487A
Date: 8/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1487/A
Problem: CF1487A
**/
func (in *CF1487A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, 101)
		for i := 0; i < n; i++ {
			num := in.NextInt()
			arr[num]++
		}

		find := false
		sol := 0
		for i := 0; i < 101; i++ {
			if find && arr[i] > 0 {
				sol += arr[i]
			}
			if arr[i] > 0 {
				find = true
			}
		}

		fmt.Println(sol)

	}
}

func NewCF1487A(r *bufio.Reader) *CF1487A {
	return &CF1487A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1487A(bufio.NewReader(os.Stdin)).Run()
}
