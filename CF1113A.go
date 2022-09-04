package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1113A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1113A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1113A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1113A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1113A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1113A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1113A
Date: 4/09/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1113/A
Problem: CF1113A
**/
func (in *CF1113A) Run() {
	n, v := in.NextInt(), in.NextInt()

	money := 0
	if v >= n {
		money = n - 1
	} else {
		diff := n - v
		money = ((diff * (diff + 1)) / 2) + (v - 1)
	}

	fmt.Println(money)

}

func NewCF1113A(r *bufio.Reader) *CF1113A {
	return &CF1113A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1113A(bufio.NewReader(os.Stdin)).Run()
}
