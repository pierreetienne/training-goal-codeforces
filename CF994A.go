package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF994A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF994A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF994A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF994A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF994A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF994A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF994A
Date: 5/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/994/A
Problem: CF994A
**/
func (in *CF994A) Run() {
	n, m := in.NextInt(), in.NextInt()
	arrN := make([]int, n)
	for i := 0; i < n; i++ {
		arrN[i] = in.NextInt()
	}

	mapa := make(map[int]int)
	for i := 0; i < m; i++ {
		mapa[in.NextInt()] = 1
	}

	var sol strings.Builder

	for i := 0; i < n; i++ {
		if mapa[arrN[i]] == 1 {
			sol.WriteString(fmt.Sprintf("%d ", arrN[i]))
		}
	}
	fmt.Println(sol.String())
}

func NewCF994A(r *bufio.Reader) *CF994A {
	return &CF994A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF994A(bufio.NewReader(os.Stdin)).Run()
}
