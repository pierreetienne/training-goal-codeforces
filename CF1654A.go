package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1654A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1654A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1654A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1654A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1654A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1654A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1654A
Date: 16/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1654/A
Problem: CF1654A
**/
func (in *CF1654A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i:=0;i<n;i++{
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		fmt.Println(arr[n-1]+arr[n-2])
	}
}

func NewCF1654A(r *bufio.Reader) *CF1654A {
	return &CF1654A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1654A(bufio.NewReader(os.Stdin)).Run()
}
