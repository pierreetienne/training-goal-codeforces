package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1625A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1625A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1625A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1625A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1625A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1625A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1625A
Date: 2/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1625/A
Problem: CF1625A
**/
func (in *CF1625A) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n, l := in.NextInt(), in.NextInt()
		ans := ""
		arr := make([]int, l)
		max := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			for j := 0; j < l; j++ {
				if (val & (1 << j)) != 0 {
					arr[j]++
				}
				if arr[j] < max {
					max = arr[j]
				}
			}
		}

		for i := l-1; i >= 0; i-- {
			if arr[i] > n/2 {
				ans+="1"
			}else {
				ans+="0"
			}
		}

		sol, _ := strconv.ParseInt(ans, 2, 32)
		fmt.Println(sol)

	}
}

func NewCF1625A(r *bufio.Reader) *CF1625A {
	return &CF1625A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1625A(bufio.NewReader(os.Stdin)).Run()
}
