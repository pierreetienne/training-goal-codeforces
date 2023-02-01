package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF818A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF818A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF818A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF818A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF818A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF818A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF818A
Date: 22/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/818/A
Problem: CF818A
**/
func (in *CF818A) Run() {
	n, k := in.NextInt64(), in.NextInt64()
	fmt.Println(n*(k+1), k*(n*(k+1)), n-(n*(k+1))-(k*(n*(k+1))))
	/*start := int64(0)
	end := n
	max := int64(0)
	for start < end {
		mid := (start + end) >> 1
		if (mid*(k+1))/k <= n/2 {
			if max < mid/k {
				max = mid / k
			}
			start = mid + 1
		} else {
			end = mid
		}
	}

	fmt.Println(max, k*max, n-(max+(k*max)))
	*/
	/*
	   				   certi = k*diplomas
	   			      winner <= n / 2
	   	              certi + k*diplimas <= n/2

	                   max(winner) => max(certi + k*deplo)


	*/
}

func NewCF818A(r *bufio.Reader) *CF818A {
	return &CF818A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF818A(bufio.NewReader(os.Stdin)).Run()
}
