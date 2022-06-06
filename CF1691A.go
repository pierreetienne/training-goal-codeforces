package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1691A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1691A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1691A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1691A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1691A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1691A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1691A
Date: 4/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1691/A
Problem: CF1691A
**/
func (in *CF1691A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		even := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i]%2 == 0 {
				even++
			}
		}

		ans := 0
		if even < n-even {
			ans = even
		}else {
			ans = n -even
		}

		fmt.Println(ans)

	}
}

func NewCF1691A(r *bufio.Reader) *CF1691A {
	return &CF1691A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1691A(bufio.NewReader(os.Stdin)).Run()
}
