package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1739B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1739B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1739B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1739B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1739B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1739B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1739B
Date: 9/30/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1739/B
Problem: CF1739B
**/
func (in *CF1739B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sol := make([]int, n)
		sol[0] = arr[0]
		ans := true
		for i := 1; i < n; i++ {
			a := sol[i-1] + arr[i]
			b := sol[i-1] - arr[i]
			if a >= 0 && b >= 0 && a != b {
				ans = false
				break
			} else {
				if a >= 0 {
					sol[i] = a
				} else {
					sol[i] = b
				}
			}
		}

		if ans {
			for i := 0; i < n; i++ {
				fmt.Print(sol[i], " ")
			}
			fmt.Println()
		} else {
			fmt.Println(-1)
		}
	}
}

func NewCF1739B(r *bufio.Reader) *CF1739B {
	return &CF1739B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1739B(bufio.NewReader(os.Stdin)).Run()
}
