package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF821A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF821A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF821A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF821A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF821A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF821A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF821A
Date: 1/7/2023
User: wotan
URL: https://codeforces.com/problemset/problem/821/A
Problem: CF821A
**/
func (in *CF821A) Run() {
	n := in.NextInt()
	m := make([][]int, n)
	rows := make([]map[int]bool, n)
	cols := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = in.NextInt()
			if rows[i] == nil {
				rows[i] = make(map[int]bool)
			}
			rows[i][m[i][j]] = true
			if cols[j] == nil {
				cols[j] = make(map[int]bool)
			}
			cols[j][m[i][j]] = true
		}
	}

	ans := true
	for i := 0; i < n && ans; i++ {
		for j := 0; j < n && ans; j++ {
			if m[i][j] != 1 {
				found := false
				for k, _ := range rows[i] {
					if _, e := cols[j][m[i][j]-k]; e {
						found = true
						break
					}

				}
				if !found {
					ans = false
				}
			}

		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No ")
	}

}

func NewCF821A(r *bufio.Reader) *CF821A {
	return &CF821A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF821A(bufio.NewReader(os.Stdin)).Run()
}
