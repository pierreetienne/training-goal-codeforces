package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1569A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1569A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1569A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1569A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1569A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1569A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1569A
Date: 18/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1569/A
Problem: CF1569A
**/
func (in *CF1569A) Run() {

	t := in.NextInt()
	for ;t > 0; t-- {
		in.NextInt()
		str := in.NextString()

		index := strings.Index(str, "ab")
		if index == -1 {
			index = strings.Index(str, "ba")
		}

		if index == -1 {
			fmt.Println(-1,-1)
		}else {
			fmt.Println(index+1, index+2)
		}
	}
}

func NewCF1569A(r *bufio.Reader) *CF1569A {
	return &CF1569A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1569A(bufio.NewReader(os.Stdin)).Run()
}
