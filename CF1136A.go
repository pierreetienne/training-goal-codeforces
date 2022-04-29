package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1136A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1136A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1136A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1136A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1136A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1136A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1136A
Date: 27/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1136/A
Problem: CF1136A
**/
func (in *CF1136A) Run() {
	n := in.NextInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = []int{in.NextInt(), in.NextInt()}
	}

	k := in.NextInt()

	for i := 0; i < n; i++ {
		if arr[i][0] <= k && k <= arr[i][1] {
			fmt.Println(n - i)
			break
		}
	}

}

func NewCF1136A(r *bufio.Reader) *CF1136A {
	return &CF1136A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1136A(bufio.NewReader(os.Stdin)).Run()
}
