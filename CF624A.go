package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF624A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF624A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF624A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF624A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF624A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF624A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF624A
Date: 12/31/2022
User: wotan
URL: https://codeforces.com/problemset/problem/624/A
Problem: CF624A
**/
func (in *CF624A) Run() {
	d, L, v1, v2 := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
	sol := float64(L-d) / float64(v2+v1)
	fmt.Printf("%.9f\n", sol)
}

func NewCF624A(r *bufio.Reader) *CF624A {
	return &CF624A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF624A(bufio.NewReader(os.Stdin)).Run()
}
