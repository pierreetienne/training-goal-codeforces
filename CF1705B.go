package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1705B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1705B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1705B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1705B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1705B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1705B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1705B
Date: 24/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1705/B
Problem: CF1705B
**/
func (in *CF1705B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		ans := int64(0)
		for i := 0; i < n; i++ {
			val := in.NextInt64()
			if i+1 < n {
				if val > 0 {
					ans += val
				}
				if val == 0 && ans > 0 {
					ans++
				}
			}

		}
		fmt.Println(ans)
	}
}

func NewCF1705B(r *bufio.Reader) *CF1705B {
	return &CF1705B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1705B(bufio.NewReader(os.Stdin)).Run()
}
