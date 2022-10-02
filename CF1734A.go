package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1734A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1734A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1734A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1734A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1734A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1734A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1734A
Date: 9/26/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1734/A
Problem: CF1734A
**/
func (in *CF1734A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		min := 100000000000
		for i := 2; i < n; i++ {
			diff := int(math.Abs(float64(arr[i-2])-float64(arr[i-1]))) + int(math.Abs(float64(arr[i])-float64(arr[i-1])))
			if diff < min {
				min = diff
			}
		}
		fmt.Println(min)
	}
}

func NewCF1734A(r *bufio.Reader) *CF1734A {
	return &CF1734A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1734A(bufio.NewReader(os.Stdin)).Run()
}
