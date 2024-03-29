package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1679A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1679A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1679A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1679A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1679A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1679A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1679A
Date: 8/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1679/A
Problem: CF1679A
**/
func (in *CF1679A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		min := int64(0)
		max := int64(0)
		a := n / 4
		llantasSobrantes := n - (a * 4)
		if llantasSobrantes == 2 || llantasSobrantes == 0 {
			max = a
			min = a
		} else {
			max = 0
		}

		b := n / 6
		llantasSobrantes = n - (b * 6)
		if llantasSobrantes == 2 || llantasSobrantes == 4 {
			min = b + 1
		} else if llantasSobrantes == 0 {
			min = b
		}

		if min == 0 || max == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(min, max)
		}

	}
}

func NewCF1679A(r *bufio.Reader) *CF1679A {
	return &CF1679A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1679A(bufio.NewReader(os.Stdin)).Run()
}
