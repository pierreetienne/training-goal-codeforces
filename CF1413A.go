package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1413A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1413A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1413A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1413A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1413A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1413A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1413A
Date: 10/31/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1413/A
Problem: CF1413A
**/
func (in *CF1413A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		sum := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			sum += arr[i]
		}

		for i := 0; i < n; i += 2 {
			fmt.Print(-arr[i+1], arr[i], " ")
		}
		fmt.Println()

	}
}

func NewCF1413A(r *bufio.Reader) *CF1413A {
	return &CF1413A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1413A(bufio.NewReader(os.Stdin)).Run()
}
