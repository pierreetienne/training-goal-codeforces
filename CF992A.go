package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF992A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF992A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF992A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF992A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF992A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF992A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF992A
Date: 1/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/992/A
Problem: CF992A
**/
func (in *CF992A) Run() {
	n := in.NextInt()
	mapa := make(map[int]int)
	for i := 0; i < n; i++ {
		x := in.NextInt()
		if x != 0 {
			mapa[x] = 1
		}
	}
	fmt.Println(len(mapa))
}

func NewCF992A(r *bufio.Reader) *CF992A {
	return &CF992A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF992A(bufio.NewReader(os.Stdin)).Run()
}
