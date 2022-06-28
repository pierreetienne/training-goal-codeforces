package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1700A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1700A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1700A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1700A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1700A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1700A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1700A
Date: 26/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1700/A
Problem: CF1700A
**/
func (in *CF1700A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt64(), in.NextInt64()
		sum := (m * (m + 1)) / 2

		for i := int64(1); i < n; i++ {
			sum += m*(i+1)
		}

		fmt.Println(sum)
	}
}

func NewCF1700A(r *bufio.Reader) *CF1700A {
	return &CF1700A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1700A(bufio.NewReader(os.Stdin)).Run()
}
