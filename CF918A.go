package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF918A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF918A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF918A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF918A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF918A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF918A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF918A
Date: 10/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/918/A
Problem: CF918A
**/
func (in *CF918A) Run() {
	n := in.NextInt()
	mapa := make(map[int]int)
	arr := make([]int, n+2)
	arr[0] = 1
	arr[1] = 1
	mapa[1] = 1
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
		mapa[arr[i]] = 1
	}

	var str strings.Builder
	for i := 0; i < n; i++ {
		if mapa[i+1] == 1 {
			str.WriteString("O")
		} else {
			str.WriteString("o")
		}
	}

	fmt.Println(str.String())
}

func NewCF918A(r *bufio.Reader) *CF918A {
	return &CF918A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF918A(bufio.NewReader(os.Stdin)).Run()
}
