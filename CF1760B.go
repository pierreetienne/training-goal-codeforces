package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1760B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1760B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1760B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1760B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1760B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1760B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1760B
Date: 11/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1760/B
Problem: CF1760B
**/
func (in *CF1760B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		max := 0
		for i := 0; i < n; i++ {
			v := int(str[i] - 'a')
			if v > max {
				max = v
			}
		}
		fmt.Println(max + 1)

	}
}

func NewCF1760B(r *bufio.Reader) *CF1760B {
	return &CF1760B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1760B(bufio.NewReader(os.Stdin)).Run()
}
