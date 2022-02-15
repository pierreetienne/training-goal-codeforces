package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF770A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF770A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF770A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF770A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF770A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF770A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF770A
Date: 14/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/770/A
Problem: CF770A
**/
func (in *CF770A) Run() {
	n, k := in.NextInt(), in.NextInt()
	str := ""

	for i := 0; i < n; i++ {
		str += string(rune('a' + (i % k)))
	}
	fmt.Println(str)
}

func NewCF770A(r *bufio.Reader) *CF770A {
	return &CF770A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF770A(bufio.NewReader(os.Stdin)).Run()
}
