package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1720B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1720B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1720B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1720B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1720B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1720B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1720B
Date: 8/24/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1720/B
Problem: CF1720B
**/
func (in *CF1720B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, 4)
		arr[0] = in.NextInt()
		arr[1] = in.NextInt()
		arr[2] = in.NextInt()
		arr[3] = in.NextInt()
		sort.Ints(arr)
		for i := 4; i < n; i++ {
			val := in.NextInt()
			if val < arr[1] {
				arr[1] = val
			} else if val < arr[0] {
				arr[0] = val
			}

			if val > arr[2] {
				arr[2] = val
			} else if val > arr[3] {
				arr[3] = val
			}
			sort.Ints(arr)
		}
		fmt.Println(arr[3] + arr[2] - arr[1] - arr[0])
	}
}

func NewCF1720B(r *bufio.Reader) *CF1720B {
	return &CF1720B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1720B(bufio.NewReader(os.Stdin)).Run()
}
