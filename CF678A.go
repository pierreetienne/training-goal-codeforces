package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF678A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF678A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF678A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF678A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF678A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF678A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF678A
Date: 9/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/678/A
Problem: CF678A
**/
func (in *CF678A) Run() {
	n, k := in.NextInt()+1, in.NextInt()
	if n%k == 0 {
		fmt.Println(n)
	} else {
		m := (k - (n % k)) + n
		fmt.Println(m)
	}

}

func NewCF678A(r *bufio.Reader) *CF678A {
	return &CF678A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF678A(bufio.NewReader(os.Stdin)).Run()
}
