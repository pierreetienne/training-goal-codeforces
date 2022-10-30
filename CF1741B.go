package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1741B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1741B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1741B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1741B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1741B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1741B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1741B
Date: 10/28/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1741/B
Problem: CF1741B
**/
func (in *CF1741B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 3 {
			fmt.Println(-1)
		} else {
			for i := 3; i <= n; i++ {
				fmt.Print(i, " ")
			}
			fmt.Println(2, 1)
		}
	}
}

func NewCF1741B(r *bufio.Reader) *CF1741B {
	return &CF1741B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1741B(bufio.NewReader(os.Stdin)).Run()
}
