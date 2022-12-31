package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1738A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1738A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1738A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1738A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1738A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1738A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1738A
Date: 10/4/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1738/A
Problem: CF1738A
**/
func (in *CF1738A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, 2)
			a[i][0] = in.NextInt()
		}
		for i := 0; i < n; i++ {
			a[i][1] = in.NextInt()
		}

		sort.Slice(a, func(i, j int) bool {
			if a[i][1] < a[j][1] {
				return true
			}
			return false
		})
		zeros := make([][]int, 0)
		ones := make([][]int, 0)
		for i := 0; i < n; i++ {
			if a[i][0] == 0 {
				zeros = append(zeros, a[i])
			} else {
				ones = append(ones, a[i])
			}
		}

		sum := 0
		if len(zeros) == len(ones) {
			if zeros[0][1] < ones[0][1] {
				sum = zeros[0][1]
				for i := 1; i < len(zeros); i++ {
					sum += zeros[i][1] * 2
				}
				for i := 0; i < len(ones); i++ {
					sum += ones[i][1] * 2
				}
			} else {
				sum = ones[0][1]
				for i := 0; i < len(zeros); i++ {
					sum += zeros[i][1] * 2
				}
				for i := 1; i < len(ones); i++ {
					sum += ones[i][1] * 2
				}
			}
		} else if len(zeros) > len(ones) {
			for i := 0; i < len(zeros)-len(ones); i++ {
				sum += zeros[i][1]
			}
			for i := len(zeros) - len(ones); i < len(zeros); i++ {
				sum += zeros[i][1] * 2
			}
			for i := 0; i < len(ones); i++ {
				sum += ones[i][1] * 2
			}
		} else {
			for i := 0; i < len(ones)-len(zeros); i++ {
				sum += ones[i][1]
			}
			for i := len(ones) - len(zeros); i < len(ones); i++ {
				sum += ones[i][1] * 2
			}
			for i := 0; i < len(zeros); i++ {
				sum += zeros[i][1] * 2
			}
		}
		fmt.Println(sum)
	}
}

func NewCF1738A(r *bufio.Reader) *CF1738A {
	return &CF1738A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1738A(bufio.NewReader(os.Stdin)).Run()
}
