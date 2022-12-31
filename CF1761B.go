package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1761B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1761B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1761B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1761B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1761B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1761B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1761B
Date: 11/24/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1761/B
Problem: CF1761B
**/
func (in *CF1761B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[int]int)
		for i := 0; i < n; i++ {
			val := in.NextInt()
			mapa[val]++

		}
		if n == 1 {
			fmt.Println(1)
		} else {

			ans := n
			if len(mapa) == 2 {
				ans = (n / 2) + 1
			}
			fmt.Println(ans)
		}
	}
}

func NewCF1761B(r *bufio.Reader) *CF1761B {
	return &CF1761B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1761B(bufio.NewReader(os.Stdin)).Run()
}
