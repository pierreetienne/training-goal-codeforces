package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1567A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1567A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1567A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1567A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1567A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1567A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1567A
Date: 19/10/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1567/A
Problem: CF1567A
**/
func (in *CF1567A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		sol := ""
		for i := 0; i < n; i++ {
			if str[i] == 'L' || str[i] == 'R' {
				sol += string(str[i])
			} else if str[i] == 'U'{
				sol += "D"
			} else {
				sol += "U"
			}
		}
		fmt.Println(sol)
	}
}

func NewCF1567A(r *bufio.Reader) *CF1567A {
	return &CF1567A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1567A(bufio.NewReader(os.Stdin)).Run()
}
