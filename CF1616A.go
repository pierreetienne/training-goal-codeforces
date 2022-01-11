package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1616A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1616A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1616A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1616A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1616A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1616A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1616A
Date: 11/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1616/A
Problem: CF1616A
**/
func (in *CF1616A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[int]int)
		for i := 0; i < n; i++ {
			val := in.NextInt()
			num, _ := mapa[val]
			mapa[val] = num + 1
		}

		for k, v := range mapa {
			if v > 1 {
				val, _ := mapa[-k]
				mapa[-k] = val + 1
			}
		}

		fmt.Println(len(mapa))
	}
}

func NewCF1616A(r *bufio.Reader) *CF1616A {
	return &CF1616A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1616A(bufio.NewReader(os.Stdin)).Run()
}
