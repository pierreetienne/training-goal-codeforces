package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1515A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1515A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1515A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1515A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1515A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1515A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1515A
Date: 28/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1515/A
Problem: CF1515A
**/
func (in *CF1515A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, x := in.NextInt(), in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		sum := 0
		ws := true
		for i := 0; i < n; i++ {
			if sum+arr[i] == x {
				if i == 0 && n != i{
					arr[i]
				} else {
					ws = false
				}
			} else {
				sum += arr[i]
			}
		}

		if ws {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1515A(r *bufio.Reader) *CF1515A {
	return &CF1515A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1515A(bufio.NewReader(os.Stdin)).Run()
}
