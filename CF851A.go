package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF851A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF851A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF851A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF851A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF851A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF851A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF851A
Date: 11/28/2022
User: wotan
URL: https://codeforces.com/problemset/problem/851/A
Problem: CF851A
**/
func (in *CF851A) Run() {
	n, k, t := in.NextInt(), in.NextInt(), in.NextInt()

	if t < k {
		fmt.Println(t)
	} else if t > n {
		fmt.Println(k - (t - n))
	} else {
		fmt.Println(k)
	}

}

func NewCF851A(r *bufio.Reader) *CF851A {
	return &CF851A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF851A(bufio.NewReader(os.Stdin)).Run()
}
