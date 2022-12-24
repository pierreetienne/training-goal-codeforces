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

type CF1767B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1767B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1767B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1767B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1767B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1767B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1767B
Date: 12/23/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1767/B
Problem: CF1767B
**/
func (in *CF1767B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		aux := arr[0]
		arr[0] = -10000000
		sort.Ints(arr)
		arr[0] = aux
		for i := 1; i < n; i++ {
			if arr[0] < arr[i] {
				diff := arr[i] - arr[0]
				arr[0] += int(math.Ceil(float64(diff) / 2.))
			}
		}

		fmt.Println(arr[0])

	}
}

func NewCF1767B(r *bufio.Reader) *CF1767B {
	return &CF1767B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1767B(bufio.NewReader(os.Stdin)).Run()
}
