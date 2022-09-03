package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1722A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1722A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1722A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1722A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1722A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1722A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1722A
Date: 9/2/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1722/A
Problem: CF1722A
**/
func (in *CF1722A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		mapa := map[uint8]int{
			'T': 1,
			'i': 1,
			'm': 1,
			'u': 1,
			'r': 1,
		}
		n := in.NextInt()
		str := in.NextString()
		ans := true
		if n == len(mapa) {
			for i := 0; i < n; i++ {
				if v, e := mapa[str[i]]; e && v == 1 {
					mapa[str[i]] = 0
					continue
				} else {
					ans = false
					break
				}
			}
		} else {
			ans = false
		}
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1722A(r *bufio.Reader) *CF1722A {
	return &CF1722A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1722A(bufio.NewReader(os.Stdin)).Run()
}
