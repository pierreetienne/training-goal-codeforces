package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1686B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1686B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1686B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1686B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1686B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1686B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1686B
Date: 9/1/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1686/B
Problem: CF1686B
**/
func (in *CF1686B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		min := 0
		countMin := -1
		count := 0
		for i := n - 1; i >= 0; i-- {
			if countMin == -1 {
				countMin = 1
				min = arr[i]
				continue
			} else if arr[i] < min {
				min = arr[i]
				countMin++
			} else {
				if countMin%2 != 0 || countMin > 1 {
					count++
				}
				countMin = -1

			}
		}
		fmt.Println(count)
	}
}

func NewCF1686B(r *bufio.Reader) *CF1686B {
	return &CF1686B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1686B(bufio.NewReader(os.Stdin)).Run()
}
