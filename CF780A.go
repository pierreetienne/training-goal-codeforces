package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF780A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF780A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF780A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF780A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF780A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF780A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF780A
Date: 12/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/780/A
Problem: CF780A
**/
func (in *CF780A) Run() {
	n := in.NextInt()
	mapa := make(map[int]int)
	ans := 0
	for i := 0; i < 2*n; i++ {
		val := in.NextInt()
		if _, ok := mapa[val]; !ok {
			mapa[val] = 1
		} else {
			delete(mapa, val)
		}
		if len(mapa) > ans {
			ans = len(mapa)
		}
	}
	fmt.Println(ans)
}

func NewCF780A(r *bufio.Reader) *CF780A {
	return &CF780A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF780A(bufio.NewReader(os.Stdin)).Run()
}
