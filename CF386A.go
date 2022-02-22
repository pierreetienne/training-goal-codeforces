package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF386A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF386A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF386A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF386A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF386A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF386A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF386A
Date: 20/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/386/A
Problem: CF386A
**/
func (in *CF386A) Run() {
	n := in.NextInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = []int{in.NextInt(), i+1}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] > arr[j][0] {
			return true
		}
		return false
	})

	fmt.Println(arr[0][1], arr[1][0])

}

func NewCF386A(r *bufio.Reader) *CF386A {
	return &CF386A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF386A(bufio.NewReader(os.Stdin)).Run()
}
