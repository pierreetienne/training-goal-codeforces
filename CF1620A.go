package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1620A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1620A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1620A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1620A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1620A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1620A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1620A
Date: 17/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1620/A
Problem: CF1620A
**/
func (in *CF1620A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		n := len(str)
		ans := true
		en := 0
		ne := 0
		for i := 0; i < n; i++ {
			if str[i] == 'E' && str[(i+1)%n] == 'N' {
				en++
			}
			if str[i] == 'N' && str[(i+1)%n] == 'E' {
				ne++
			}
		}

		if en == n || ne == n {
			ans = false
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1620A(r *bufio.Reader) *CF1620A {
	return &CF1620A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1620A(bufio.NewReader(os.Stdin)).Run()
}
