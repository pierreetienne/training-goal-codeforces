package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1732B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1732B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1732B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1732B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1732B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1732B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1732B
Date: 11/18/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1732/B
Problem: CF1732B
**/
func (in *CF1732B) Run() {
	var sol strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		var last uint8 = '-'
		ans := 0

		d := 0
		for d < n {
			if str[d] == '0' {
				d++
			} else {
				break
			}
		}

		for i := n - 1; i >= d; i-- {
			if i == n-1 {
				last = str[i]
			}

			if last != str[i] {
				ans++
				last = str[i]
			}
		}

		sol.WriteString(fmt.Sprintf("%v\n", ans))
	}
	fmt.Print(sol.String())
}

func NewCF1732B(r *bufio.Reader) *CF1732B {
	return &CF1732B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1732B(bufio.NewReader(os.Stdin)).Run()
}
