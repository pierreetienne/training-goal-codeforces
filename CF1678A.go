package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1678A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1678A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1678A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1678A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1678A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1678A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1678A
Date: 6/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1678/A
Problem: CF1678A
**/
func (in *CF1678A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := map[int]int{}

		repeat := false
		for i := 0; i < n; i++ {
			num := in.NextInt()
			value, _ := mapa[num]
			mapa[num] = value + 1
			if value+1 >= 2 {
				repeat = true
			}
		}

		zero, exist := mapa[0]
		ans := 0
		if exist {
			ans = n - zero
		} else if repeat {
			ans = n
		} else {
			ans = n + 1
		}
		fmt.Println(ans)
	}
}

func NewCF1678A(r *bufio.Reader) *CF1678A {
	return &CF1678A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1678A(bufio.NewReader(os.Stdin)).Run()
}
