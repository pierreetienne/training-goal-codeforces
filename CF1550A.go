package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1550A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1550A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1550A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1550A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1550A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1550A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1550A
Date: 19/10/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1550/A
Problem: CF1550A
**/
var n = 0

func (in *CF1550A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n = in.NextInt()

		if n < 2 {
			fmt.Println(n)
		} else{
			sol := 2
			aux := 3
			for i := 4; i < n; i+=aux{
				sol++
				aux+=2
			}
			fmt.Println(sol)
		}

	}
}

func f(sum, m int) int {
	if n-sum == 0 {
		return 0
	}
	if n-sum < 0 {
		return -1
	}

	a := f(sum+m+1, m+1)
	b := f(sum+m+2, m+1)
	c := f(sum+m+1, m+2)
	d := f(sum+m+2, m+2)
	e := f(sum+1, m)

	arr := []int{a, b, c, d, e}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			//memo[sum][n] =
			return arr[i] + 1
		}
	}

	return -1
}

func NewCF1550A(r *bufio.Reader) *CF1550A {
	return &CF1550A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1550A(bufio.NewReader(os.Stdin)).Run()
}
