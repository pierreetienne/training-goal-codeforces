package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1726A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1726A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1726A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1726A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1726A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1726A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1726A
Date: 9/6/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1726/A
Problem: CF1726A
**/
func (in *CF1726A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		min, max := 9999, 0
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			val := in.NextInt()
			arr[i] = val
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}

		ok := false
		for i := 0; i < n; i++ {
			if arr[i] == max && arr[(i+1)%n] == min {
				ok = true
				break
			}
		}

		if arr[0] == min || arr[n-1] == max || ok {
			fmt.Println(max - min)
		} else {
			ans := -999999
			for i, j := 0, n-1; i < n; i++ {
				val := arr[j] - arr[i]
				if val > ans {
					ans = val
				}
				j = (j + 1) % n
			}
			fmt.Println(Max(max-arr[0], Max(arr[n-1]-min, ans)))
		}

	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewCF1726A(r *bufio.Reader) *CF1726A {
	return &CF1726A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1726A(bufio.NewReader(os.Stdin)).Run()
}
