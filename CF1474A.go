package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1474A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1474A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1474A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1474A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1474A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1474A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1474A
Date: 17/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1474/A
Problem: CF1474A
**/
func (in *CF1474A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		a := in.NextString()
		var sol strings.Builder
		result := ""
		for i := 0; i < n; i++ {
			choose := ""
			if i == 0 {
				choose = "1"
			} else {
				if a[i] == '1' {
					if result == "2" {
						choose = "0"
					} else {
						choose = "1"
					}
				} else {
					if result == "1" {
						choose = "0"
					} else {
						choose = "1"
					}
				}
			}

			if choose == "1" && a[i] == '1' {
				result = "2"
			} else if (choose == "0" && a[i] == '1') || (choose == "1" && a[i] == '0') {
				result = "1"
			} else if choose == "0" && a[i] == '0' {
				result = "0"
			}
			sol.WriteString(choose)
		}
		fmt.Println(sol.String())
	}
}

func NewCF1474A(r *bufio.Reader) *CF1474A {
	return &CF1474A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1474A(bufio.NewReader(os.Stdin)).Run()
}
