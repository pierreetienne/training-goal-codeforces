package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1742AA struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1742AA) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1742AA) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1742AA) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1742AA) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1742AA) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1742AA
Date: 10/15/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1742/A
Problem: CF1742AA
**/
func (in *CF1742AA) Run() {

}

func NewCF1742AA(r *bufio.Reader) *CF1742AA {
	return &CF1742AA{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1742AA(bufio.NewReader(os.Stdin)).Run()
}
