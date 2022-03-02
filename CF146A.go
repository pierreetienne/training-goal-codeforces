package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF146A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF146A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF146A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF146A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF146A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF146A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF146A
Date: 28/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/146/A
Problem: CF146A
**/
func (in *CF146A) Run() {
	n := in.NextInt()
	str := in.NextString()
	sumA := 0
	ans := true
	for i := 0; i < n; i++ {
		if i >= n/2 {
			sumA -= int(str[i] - '0')
		} else {
			sumA += int(str[i] - '0')
		}

		if str[i] != '4' && str[i] != '7' {
			ans = false
		}
	}

	if sumA != 0 {
		ans = false
	}
	if ans {
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}

}

func NewCF146A(r *bufio.Reader) *CF146A {
	return &CF146A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF146A(bufio.NewReader(os.Stdin)).Run()
}
