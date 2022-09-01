package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1672A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1672A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1672A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1672A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1672A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1672A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1672A
Date: 8/23/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1672/A
Problem: CF1672A
**/
func (in *CF1672A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		sum := 0
		for i := 0; i < n; i++ {
			sum += in.NextInt() - 1
		}

		turn := "errorgorn"

		if sum%2 == 0 {
			turn = "maomao90"
		}
		fmt.Println(turn)
	}
}

func NewCF1672A(r *bufio.Reader) *CF1672A {
	return &CF1672A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1672A(bufio.NewReader(os.Stdin)).Run()
}
