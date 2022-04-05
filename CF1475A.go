	package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1475A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1475A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1475A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1475A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1475A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1475A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1475A
Date: 3/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1475/A
Problem: CF1475A
**/
func (in *CF1475A) Run() {
	for t:=in.NextInt();t>0; t--{
		n := in.NextInt64()
		if n%2 == 0 {
			fmt.Println("NO")
		}else{
			fmt.Println("YES")
		}
	}
}

func NewCF1475A(r *bufio.Reader) *CF1475A {
	return &CF1475A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1475A(bufio.NewReader(os.Stdin)).Run()
}
