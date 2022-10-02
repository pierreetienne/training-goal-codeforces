package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1733A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1733A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1733A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1733A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1733A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1733A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1733A
Date: 23/09/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1733/A
Problem: CF1733A
**/
func (in *CF1733A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		maxi := int64(0)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		aux := k
		for i := k; i < len(arr) && aux > 0; i++ {
			if arr[i%k] < arr[i] {
				arr[i%k] = arr[i]
				aux++
			}
		}

		fmt.Println(maxi)
	}
}

func NewCF1733A(r *bufio.Reader) *CF1733A {
	return &CF1733A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1733A(bufio.NewReader(os.Stdin)).Run()
}
