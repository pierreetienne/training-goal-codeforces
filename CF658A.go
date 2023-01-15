package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF658A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF658A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF658A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF658A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF658A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF658A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF658A
Date: 1/14/2023
User: wotan
URL: https://codeforces.com/problemset/problem/658/A
Problem: CF658A
**/
func (in *CF658A) Run() {
	n, c := in.NextInt(), in.NextInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = in.NextInt()
	}
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = in.NextInt()
	}

	sumA := 0
	minutes := 0
	for i := 0; i < n; i++ {
		minutes += t[i]
		sumA += int(math.Max(float64(0), float64(p[i]-(c*minutes))))
	}

	sumB := 0
	minutes = 0
	for i := n - 1; i >= 0; i-- {
		minutes += t[i]
		sumB += int(math.Max(float64(0), float64(p[i]-(c*minutes))))
	}

	if sumA > sumB {
		fmt.Println("Limak")
	} else if sumB > sumA {
		fmt.Println("Radewoosh")
	} else {
		fmt.Println("Tie")
	}
}

func NewCF658A(r *bufio.Reader) *CF658A {
	return &CF658A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF658A(bufio.NewReader(os.Stdin)).Run()
}
