package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF755A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF755A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF755A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF755A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF755A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF755A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF755A
Date: 7/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/755/A
Problem: CF755A
**/
func (in *CF755A) Run() {
	n := in.NextInt()

	prime := make([]bool, 20000)

	for i := 2; i < len(prime); i++ {
		for j := 2* i; j < len(prime); j += i {
			prime[j] = true
		}
	}


	for i := 1; i < 100; i++ {
		if prime[i*n+1] {
			fmt.Println(i)
			break
		}
	}
}

func NewCF755A(r *bufio.Reader) *CF755A {
	return &CF755A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF755A(bufio.NewReader(os.Stdin)).Run()
}
