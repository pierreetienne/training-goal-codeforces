package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1686A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1686A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1686A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1686A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1686A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1686A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1686A
Date: 1/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1686/A
Problem: CF1686A
**/
func (in *CF1686A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		mapa := map[int]bool{}
		sum := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			mapa[arr[i]] = true
			sum += arr[i]
		}

		ans := "NO"

		if len(mapa) == 1 {
			ans = "YES"
		} else {
			for i := 0; i < n; i++ {
				val := sum - arr[i]
				if val%(n-1) == 0 && int(val/(n-1)) == arr[i] {
					ans = "YES"
				}

			}
		}
		fmt.Println(ans)

	}
}

func NewCF1686A(r *bufio.Reader) *CF1686A {
	return &CF1686A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1686A(bufio.NewReader(os.Stdin)).Run()
}
