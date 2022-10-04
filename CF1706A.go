package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1706A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1706A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1706A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1706A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1706A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1706A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Ez for gopherbots 60 ms 
Run solve the problem CF1706A
Date: 22/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1706/A
Problem: CF1706A
**/
func (in *CF1706A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		arr := make([]int, n)
		mapa := make(map[int]bool)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		for i := 0; i < n; i++ {
			a := arr[i]
			b := m + 1 - arr[i]
			_, eA := mapa[a]
			_, eB := mapa[b]

			if a < b && !eA {
				mapa[a] = true
			} else if b < a && !eB {
				mapa[b] = true
			} else if !eA {
				mapa[a] = true
			} else {
				mapa[b] = true
			}
		}

		var ans strings.Builder
		for i := 0; i < m; i++ {
			if _, e := mapa[i+1]; e {
				ans.WriteString("A")
			} else {
				ans.WriteString("B")
			}
		}

		fmt.Println(ans.String())
	}
}

func NewCF1706A(r *bufio.Reader) *CF1706A {
	return &CF1706A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1706A(bufio.NewReader(os.Stdin)).Run()
}
