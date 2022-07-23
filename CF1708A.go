package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1708A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1708A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1708A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1708A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1708A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1708A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1708A
Date: 21/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1708/A
Problem: CF1708A
**/
func (in *CF1708A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		ans := true
		for i := 0; i < n; i++ {
			if arr[i] % arr[0] != 0 {
				ans = false
				break
			}
		}
		if ans {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}

}

func NewCF1708A(r *bufio.Reader) *CF1708A {
	return &CF1708A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1708A(bufio.NewReader(os.Stdin)).Run()
}
