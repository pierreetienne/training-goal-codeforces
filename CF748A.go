package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF748A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF748A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF748A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF748A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF748A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF748A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF748A
Date: 1/4/2023
User: wotan
URL: https://codeforces.com/problemset/problem/748/A
Problem: CF748A
**/
func (in *CF748A) Run() {
	_, m, k := in.NextInt(), in.NextInt(), in.NextInt()
	lane := 1
	desk := 1
	for {
		if lane*(m*2) < k {
			lane++

		} else {
			for i := (lane-1)*(m*2) - 1; i <= (lane)*(m*2); i += 2 {
				if k != i && k != i+1 {
					desk++
				} else {
					break
				}
			}
			break
		}
	}

	side := 'L'
	if k%2 == 0 {
		side = 'R'
	}
	fmt.Println(lane, desk-1, string(side))
}

func NewCF748A(r *bufio.Reader) *CF748A {
	return &CF748A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF748A(bufio.NewReader(os.Stdin)).Run()
}
