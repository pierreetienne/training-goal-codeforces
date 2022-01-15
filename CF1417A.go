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

type CF1417A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1417A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1417A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1417A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1417A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1417A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1417A
Date: 15/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1417/A
Problem: CF1417A
**/
func (in *CF1417A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		sol := 0
		for i := 1; i < n; i++ {
			diff := k - arr[i]
			if diff > 0 {
				sol += int(math.Floor(float64(diff) / float64(arr[0])))
			}
		}
		fmt.Println(sol)
	}
}

func NewCF1417A(r *bufio.Reader) *CF1417A {
	return &CF1417A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1417A(bufio.NewReader(os.Stdin)).Run()
}
