package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1501A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1501A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1501A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1501A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1501A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1501A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1501A
Date: 12/7/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1501/A
Problem: CF1501A
**/
func (in *CF1501A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([][]int, n+1)
		arr[0] = []int{0, 0}
		for i := 0; i < n; i++ {
			arr[i+1] = []int{in.NextInt(), in.NextInt()}
		}
		arr[0][1] = arr[1][0]
		tt := make([]int, n)
		for i := 0; i < n; i++ {
			tt[i] = in.NextInt()
		}

		lastDepart := arr[1][0]
		lastStart := 0

		for index := 1; index <= n; index++ {

			start := lastDepart + tt[index-1] + (arr[index][0] - arr[index-1][1])
			lastStart = start
			min := int(math.Ceil(float64(arr[index][1]-arr[index][0]) / 2.))

			depart := 0
			if start+min >= arr[index][1] {
				depart = start + min
			} else {
				depart = arr[index][1]
			}

			lastDepart = depart

		}

		fmt.Println(lastStart)
	}
}

func NewCF1501A(r *bufio.Reader) *CF1501A {
	return &CF1501A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1501A(bufio.NewReader(os.Stdin)).Run()
}
