package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1681A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1681A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1681A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1681A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1681A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1681A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1681A
Date: 25/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1681/A
Problem: CF1681A
**/
func (in *CF1681A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		nums := make([]int, 51)
		for i := 0; i < n; i++ {
			nums[in.NextInt()]++
		}
		m := in.NextInt()

		b := make([]int, 51)
		for i := 0; i < m; i++ {
			b[in.NextInt()]++
		}

		x := false
		y := false
		for i := 50; i >= 0; i-- {
			if b[i] > 0 && nums[i] > 0 {
				x = true
				y = false
				break
			}
			if b[i] > 0 && nums[i] == 0 {
				y = false
				x = false
				break
			}

			if b[i] == 0 && nums[i] > 0 {
				y = true
				x = true
				break
			}
		}

		if x {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}

		if y {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}
	}
}

func NewCF1681A(r *bufio.Reader) *CF1681A {
	return &CF1681A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1681A(bufio.NewReader(os.Stdin)).Run()
}
