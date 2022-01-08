package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1411A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1411A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1411A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1411A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1411A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1411A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1411A
Date: 7/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1411/A
Problem: CF1411A
**/
func (in *CF1411A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		count := 0
		for i := n - 1; i >= 0; i-- {
			if str[i] == ')' {
				count++
			} else {
				break
			}
		}

		mid := int(math.Floor(float64(n) / 2.0))
		if count > mid {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	}
}

func NewCF1411A(r *bufio.Reader) *CF1411A {
	return &CF1411A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1411A(bufio.NewReader(os.Stdin)).Run()
}
