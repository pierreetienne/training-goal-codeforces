package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1676B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1676B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1676B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1676B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1676B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1676B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1676B
Date: 12/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1676/B
Problem: CF1676B
**/
func (in *CF1676B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		min := 999999999
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			val := in.NextInt()
			arr[i] = val
			if val < min {
				min = val
			}
		}
		ans := 0

		for i := 0; i < n; i++ {
			ans += arr[i] - min
		}
		fmt.Println(ans)
	}

}

func NewCF1676B(r *bufio.Reader) *CF1676B {
	return &CF1676B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1676B(bufio.NewReader(os.Stdin)).Run()
}
