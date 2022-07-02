package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF294A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF294A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF294A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF294A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF294A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF294A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF294A
Date: 25/01/22
User: pepradere
URL: https://codeforces.com/contest/294/problem/A
Problem: CF294A
**/
func (in *CF294A) Run() {
	a := make([]int, 101)
	n := in.NextInt()
	for i := 0; i < n; i++ {
		a[i] = in.NextInt()
	}
	m := in.NextInt()
	for i := 1; i <= m; i++ {
		x, y := in.NextInt(), in.NextInt()
		fmt.Println()
		a[x-1] += y - 1
		a[x+1] += a[x] - y
		a[x] = 0
	}
	var str strings.Builder
	for i := 1; i <=n; i++ {
		str.WriteString(fmt.Sprintf("%d\n", a[i]))
	}
	fmt.Print(str.String())

}

func NewCF294A(r *bufio.Reader) *CF294A {
	return &CF294A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF294A(bufio.NewReader(os.Stdin)).Run()
}
