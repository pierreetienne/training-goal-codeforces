package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1629A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1629A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1629A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1629A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1629A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1629A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1629A
Date: 2/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1629/A
Problem: CF1629A
**/
func (in *CF1629A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		m := make([][]int, n)
		for i := 0; i < n; i++ {
			m[i] = make([]int, 2)
			m[i][0] = in.NextInt()
		}
		for i := 0; i < n; i++ {
			m[i][1] = in.NextInt()
		}

		sort.Slice(m, func(i, j int) bool {
			if m[i][0]-m[j][0] == 0 {
				if m[i][1]-m[j][1] > 0 {
					return false
				}
				return true
			}
			if m[i][0]-m[j][0] > 0 {
				return false
			}
			return true
		})

		result := k
		for i := 0; i < n; i++ {
			if m[i][0] <= result {
				result += m[i][1]
			}
		}
		fmt.Println(result)
	}
}

func NewCF1629A(r *bufio.Reader) *CF1629A {
	return &CF1629A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1629A(bufio.NewReader(os.Stdin)).Run()
}
