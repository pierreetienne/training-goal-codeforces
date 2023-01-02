package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Test123 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *Test123) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *Test123) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *Test123) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *Test123) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *Test123) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem Test123
Date: 11/30/2022
User: wotan
URL:
Problem: Test123
**/
func (in *Test123) Run() {
	n := in.NextInt()
	sol := ""
	for i := 0; i < n; i++ {
		sol += fmt.Sprintf("%d\n", solve3(in.NextInt(), in))
	}
	fmt.Println(sol)
}

func solve3(n int, in *Test123) int {
	countOnes := 0
	base := 0
	indexFirstZero := -1
	indexLastOne := -1
	for i := 0; i < n; i++ {
		s := in.NextInt()
		if s == 0 {
			if indexFirstZero == -1 {
				indexFirstZero = i
			}
			base = base + countOnes
		} else {
			indexLastOne = i
			countOnes++
		}
	}

	return max(max(base, base-indexFirstZero+(n-countOnes-1)), base+(countOnes-1)-(n-(indexLastOne+1)))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewTest123(r *bufio.Reader) *Test123 {
	return &Test123{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewTest123(bufio.NewReader(os.Stdin)).Run()
}
