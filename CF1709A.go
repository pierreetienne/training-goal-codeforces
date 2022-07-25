package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1709A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1709A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1709A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1709A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1709A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1709A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1709A
Date: 23/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1709/A
Problem: CF1709A
**/
func (in *CF1709A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		key := in.NextInt() - 1
		arr := []int{
			in.NextInt() - 1,
			in.NextInt() - 1,
			in.NextInt() - 1,
		}

		ans := true
		//|| arr[arr[arr[key]]] == -1
		if arr[key] == -1 || arr[arr[key]] == -1 {
			ans = false
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1709A(r *bufio.Reader) *CF1709A {
	return &CF1709A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1709A(bufio.NewReader(os.Stdin)).Run()
}
