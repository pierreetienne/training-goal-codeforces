package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1742B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1742B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1742B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1742B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1742B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1742B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1742B
Date: 10/14/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1742/B
Problem: CF1742B
**/
func (in *CF1742B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[int]int)
		ans := true
		for i := 0; i < n; i++ {
			val := in.NextInt()
			mapa[val]++
			if v := mapa[val]; v > 1 {
				ans = false
			}
		}
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1742B(r *bufio.Reader) *CF1742B {
	return &CF1742B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1742B(bufio.NewReader(os.Stdin)).Run()
}
