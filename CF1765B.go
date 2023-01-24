package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1765B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1765B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1765B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1765B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1765B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1765B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1765B
Date: 18/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1765/B
Problem: CF1765B
**/
func (in *CF1765B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		f := true
		ans := true
		for i := 0; i < n-1; i++ {
			if !f {
				if str[i] != str[i+1] {
					ans = false
				}
				i++
			}
			f = !f
		}

		if ans && f {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1765B(r *bufio.Reader) *CF1765B {
	return &CF1765B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1765B(bufio.NewReader(os.Stdin)).Run()
}
