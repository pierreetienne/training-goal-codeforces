package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1627A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1627A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1627A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1627A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1627A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1627A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1627A
Date: 18/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1627/A
Problem: CF1627A
**/
func (in *CF1627A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m, r, c := in.NextInt(), in.NextInt(), in.NextInt()-1, in.NextInt()-1
		arr := make([]string, n)
		containsB := false
		for i := 0; i < n; i++ {
			arr[i] = in.NextString()
			if strings.Contains(arr[i], "B") {
				containsB = true
			}
		}

		if !containsB {
			fmt.Println(-1)
		} else {
			ans := 2
			for i := 0; i < n; i++ {
				if arr[i][c] == 'B' {
					ans = 1
				}
			}
			for i := 0; i < m; i++ {
				if arr[r][i] == 'B' {
					ans = 1
				}
			}
			if arr[r][c] == 'B' {
				ans = 0
			}
			fmt.Println(ans)
		}
	}

}

func NewCF1627A(r *bufio.Reader) *CF1627A {
	return &CF1627A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1627A(bufio.NewReader(os.Stdin)).Run()
}
