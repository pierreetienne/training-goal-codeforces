package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1031A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1031A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1031A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1031A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1031A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1031A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1031A
Date: 8/12/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1031/A
Problem: CF1031A
**/
func (in *CF1031A) Run() {
	w, h, k := in.NextInt(), in.NextInt(), in.NextInt()

	ans := 0
	for k > 0 {

		ans += (w + w + h + h) - 4
		w -= 4
		h -= 4
		k--
		if w <= 0 || h <= 0 {
			break
		}
	}

	fmt.Println(ans)
}

func NewCF1031A(r *bufio.Reader) *CF1031A {
	return &CF1031A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1031A(bufio.NewReader(os.Stdin)).Run()
}
