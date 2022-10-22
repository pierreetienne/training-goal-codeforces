package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1742C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1742C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1742C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1742C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1742C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1742C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots
Run solve the problem CF1742C
Date: 10/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1742/C
Problem: CF1742C
**/
func (in *CF1742C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		m := make([]string, 8)
		in.GetLine()
		for i := 0; i < 8; i++ {
			m[i] = in.NextString()
		}

		ans := uint8(0)
		for i := 0; i < 8; i++ {
			r := 0
			b := 0

			for j := 0; j < 8; j++ {
				if m[i][j] == 'R' {
					r++
				}
			}
			if r == 8 {
				ans = 'R'
				break
			}

			for j := 0; j < 8; j++ {
				if m[j][i] == 'B' {
					b++
				}
			}

			if b == 8 {
				ans = 'B'
				break
			}

		}

		fmt.Println(string(rune(ans)))
	}
}

func NewCF1742C(r *bufio.Reader) *CF1742C {
	return &CF1742C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1742C(bufio.NewReader(os.Stdin)).Run()
}
