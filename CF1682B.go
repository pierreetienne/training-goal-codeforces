package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1682B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1682B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1682B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1682B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1682B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1682B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1682B
Date: 30/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1681/B
Problem: CF1682B
**/
func (in *CF1682B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}
		m := in.NextInt()
		top := 0
		for i := 0; i < m; i++ {
			bi := in.NextInt()
			top = ((bi % n) + (top % n)) % n
		}
		fmt.Println(a[(top)%n])
	}
}

func NewCF1682B(r *bufio.Reader) *CF1682B {
	return &CF1682B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1682B(bufio.NewReader(os.Stdin)).Run()
}
