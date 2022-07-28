package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1672B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1672B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1672B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1672B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1672B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1672B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1672B
Date: 27/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1672/B
Problem: CF1672B
**/
func (in *CF1672B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ans := true

		if len(str) < 2 {
			ans = false
		}

		b := 0
		a := 0
		for i := 0; i < len(str) && ans; i++ {
			if str[i] == 'B' {
				b++
			} else {
				a++
			}
			if b > a {
				ans = false
			}
		}

		if b > a {
			ans = false
		}

		if ans && str[0] == 'B' {
			ans = false
		}

		if ans && str[len(str)-1] == 'A' {
			ans = false
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func NewCF1672B(r *bufio.Reader) *CF1672B {
	return &CF1672B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1672B(bufio.NewReader(os.Stdin)).Run()
}
