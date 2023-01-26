package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1056A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1056A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1056A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1056A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1056A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1056A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1056A
Date: 1/23/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1056/A
Problem: CF1056A
**/
func (in *CF1056A) Run() {
	mapa := make(map[int]int)
	var sol strings.Builder
	n := in.NextInt()
	for j := 0; j < n; j++ {
		r := in.NextInt()
		for i := 0; i < r; i++ {
			val := in.NextInt()
			mapa[val]++
			if c, _ := mapa[val]; c == n {
				sol.WriteString(fmt.Sprintf("%d ", val))
			}
		}
	}
	fmt.Println(sol.String())

}

func NewCF1056A(r *bufio.Reader) *CF1056A {
	return &CF1056A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1056A(bufio.NewReader(os.Stdin)).Run()
}
