package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF177A1 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF177A1) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF177A1) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF177A1) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF177A1) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF177A1) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF177A1
Date: 1/13/2023
User: wotan
URL: https://codeforces.com/problemset/problem/177/A1
Problem: CF177A1
**/
func (in *CF177A1) Run() {
	n := in.NextInt()
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = in.NextInt()
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += m[i][i]
		m[i][i] = 0
		sum += m[n-i-1][i]
		m[n-i-1][i] = 0
	}

	for i := 0; i < n; i++ {
		sum += m[i][(n-1)/2]

		m[i][(n-1)/2] = 0
		sum += m[(n-1)/2][i]
		m[(n-1)/2][i] = 0
	}

	fmt.Println(sum)
}

func NewCF177A1(r *bufio.Reader) *CF177A1 {
	return &CF177A1{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF177A1(bufio.NewReader(os.Stdin)).Run()
}
