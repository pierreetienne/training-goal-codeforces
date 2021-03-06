package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1270A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1270A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1270A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1270A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1270A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1270A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1270A
Date: 23/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1270/A
Problem: CF1270A
**/
func (in *CF1270A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		_, k1, k2 := in.NextInt(), in.NextInt(), in.NextInt()

		max1 := 0
		max2 := 0
		for i := 0; i < k1; i++ {
			val := in.NextInt()
			if val > max1 {
				max1 = val
			}
		}
		for i := 0; i < k2; i++ {
			val := in.NextInt()
			if val > max2 {
				max2 = val
			}
		}


		if max1  > max2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1270A(r *bufio.Reader) *CF1270A {
	return &CF1270A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1270A(bufio.NewReader(os.Stdin)).Run()
}
