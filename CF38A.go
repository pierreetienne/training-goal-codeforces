package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF38A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF38A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF38A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF38A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF38A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF38A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF38A
Date: 4/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/38/A
Problem: CF38A
**/
func (in *CF38A) Run() {
	n := in.NextInt()

	arr := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		arr[i] = in.NextInt()
	}

	a, b := in.NextInt(), in.NextInt()

	sum := 0
	for i := 0; i < n-1; i++ {
		from := i + 1
		if from >= a && from < b {
			sum += arr[i]
		}
	}

	fmt.Println(sum)
}

func NewCF38A(r *bufio.Reader) *CF38A {
	return &CF38A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF38A(bufio.NewReader(os.Stdin)).Run()
}
