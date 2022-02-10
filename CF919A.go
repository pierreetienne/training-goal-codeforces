package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF919A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF919A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF919A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF919A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF919A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF919A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF919A
Date: 9/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/919/A
Problem: CF919A
**/
func (in *CF919A) Run() {
	n, m := in.NextInt(), in.NextInt()
	ans := float64(100000000)
	for i := 0; i < n; i++ {
		a, b := in.NextInt(), in.NextInt()
		res := (float64(m) * float64(a)) / float64(b)
		if res < ans {
			ans = res
		}
	}
	fmt.Printf("%.6f\n", ans)
}

func NewCF919A(r *bufio.Reader) *CF919A {
	return &CF919A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF919A(bufio.NewReader(os.Stdin)).Run()
}
