package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1760D struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1760D) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1760D) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1760D) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1760D) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1760D) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1760D
Date: 11/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1760/D
Problem: CF1760D
**/
func (in *CF1760D) Run() {
	for t := in.NextInt(); t > 0; t-- {
		nL := in.NextInt()
		arr := make([]int, 0)
		last := -1

		for i := 0; i < nL; i++ {
			val := in.NextInt()
			if last != val {
				arr = append(arr, val)
			}
			last = val
		}
		n := len(arr)

		valley := 0
		for i := 0; i < n; i++ {
			if i == 0 && i+1 < n {
				if arr[i+1] > arr[i] {
					valley++
				}
			} else if i == n-1 && i-1 >= 0 {
				if arr[i-1] > arr[i] {
					valley++
				}
			} else {
				if i-1 >= 0 && i < n && arr[i-1] > arr[i] && arr[i+1] > arr[i] {
					valley++
				}
			}
		}

		if valley == 1 || len(arr) == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1760D(r *bufio.Reader) *CF1760D {
	return &CF1760D{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1760D(bufio.NewReader(os.Stdin)).Run()
}
