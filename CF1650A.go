package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1650A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1650A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1650A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1650A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1650A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1650A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1650A
Date: 11/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1650/A
Problem: CF1650A
**/
func (in *CF1650A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		char := in.NextString()[0]

		ans := false
		for i := 0; i < len(str); i++ {
			if str[i] == char && i%2 == 0{
					ans = true
					break
			}
		}

		if ans {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}

func NewCF1650A(r *bufio.Reader) *CF1650A {
	return &CF1650A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1650A(bufio.NewReader(os.Stdin)).Run()
}
