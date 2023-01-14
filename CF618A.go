package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF618A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF618A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF618A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF618A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF618A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF618A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF618A
Date: 8/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/618/A
Problem: CF618A
**/
func (in *CF618A) Run() {
	n := in.NextInt()
	var str strings.Builder
	for i := 32; i >= 0; i-- {
		if n&(1<<i) != 0 {
			str.WriteString(fmt.Sprintf("%v ", i+1))
		}
	}
	fmt.Println(str.String())
}

func NewCF618A(r *bufio.Reader) *CF618A {
	return &CF618A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF618A(bufio.NewReader(os.Stdin)).Run()
}
