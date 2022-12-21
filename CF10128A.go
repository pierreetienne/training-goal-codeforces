package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF10128A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF10128A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF10128A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF10128A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF10128A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF10128A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF10128A
Date: 10/12/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1028/A
Problem: CF10128A
**/
func (in *CF10128A) Run() {
	n, m := in.NextInt(), in.NextInt()
	str := make([]string, n)
	b := 0
	indexI, indexJ := -1, -1
	for i := 0; i < n; i++ {
		str[i] = in.NextString()
		s := 0
		for j := 0; j < m; j++ {
			if str[i][j] == 'B' {
				if indexI == -1 {
					indexI = i
					indexJ = j
				}
				s++
			}
		}
		if s > 0 {
			b = s
		}
	}

	if b == 1 {
		fmt.Println(indexI+1, indexJ+1)
	} else {
		mid := (b + 1) / 2
		fmt.Println(indexI+mid, indexJ+mid)
	}
}

func NewCF10128A(r *bufio.Reader) *CF10128A {
	return &CF10128A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF10128A(bufio.NewReader(os.Stdin)).Run()
}
