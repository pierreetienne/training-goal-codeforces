package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1017A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1017A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1017A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1017A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1017A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1017A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1017A
Date: 9/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1017/A
Problem: CF1017A
**/
func (in *CF1017A) Run() {
	n := in.NextInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
	}
	a := 0
	pos := 1
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < 4; j++ {
			num := in.NextInt()
			if i == 0 {
				a += num
			}
			sum += num
		}
		if sum > a {
			pos++
		}

	}
	fmt.Println(pos)
}

func NewCF1017A(r *bufio.Reader) *CF1017A {
	return &CF1017A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1017A(bufio.NewReader(os.Stdin)).Run()
}
