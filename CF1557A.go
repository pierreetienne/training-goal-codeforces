package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1557A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1557A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1557A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1557A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1557A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1557A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1557A
Date: 6/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1557/A
Problem: CF1557A
**/
func (in *CF1557A) Run() {

	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		sum := float64(0)
		for i := 0; i < n-1; i++ {
			sum += float64(arr[i])
		}
		sol := float64(arr[n-1]) + (sum / float64(n-1))
		fmt.Printf("%.7f\n", sol)

	}
}

func NewCF1557A(r *bufio.Reader) *CF1557A {
	return &CF1557A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1557A(bufio.NewReader(os.Stdin)).Run()
}
