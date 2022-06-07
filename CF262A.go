package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF262A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF262A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF262A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF262A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF262A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF262A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF262A
Date: 5/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/262/A
Problem: CF262A
**/
func (in *CF262A) Run() {
	n, k := in.NextInt(), in.NextInt()
	ans := 0
	for i := 0; i < n; i++ {
		c := in.NextString()
		count :=0
		for j := 0; j < len(c); j++ {
			if c[j]=='7' || c[j]=='4' {
				count++
			}
		}
		if count <= k {
			ans++
		}
	}
	fmt.Println(ans)
}

func NewCF262A(r *bufio.Reader) *CF262A {
	return &CF262A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF262A(bufio.NewReader(os.Stdin)).Run()
}
