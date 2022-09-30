package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1734B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1734B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1734B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1734B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1734B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1734B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1734B
Date: 9/28/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1734/B
Problem: CF1734B
**/
func (in *CF1734B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					fmt.Print(1, " ")
				} else {
					fmt.Print(0, " ")
				}
			}
			fmt.Println()
		}
	}

}

func NewCF1734B(r *bufio.Reader) *CF1734B {
	return &CF1734B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1734B(bufio.NewReader(os.Stdin)).Run()
}
