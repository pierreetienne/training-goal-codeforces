package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1743B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1743B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1743B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1743B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1743B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1743B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1743B
Date: 23/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1743/B
Problem: CF1743B
**/
func (in *CF1743B) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		for i, j := 1, n; i <= j; i++ {
			val := 0
			if i == j {
				val = i
			} else {
				val = j
				ans.WriteString(fmt.Sprintf("%v ", i))
			}
			ans.WriteString(fmt.Sprintf("%v ", val))
			j--
		}
		ans.WriteString("\n")
	}
	fmt.Print(ans.String())
}

func NewCF1743B(r *bufio.Reader) *CF1743B {
	return &CF1743B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1743B(bufio.NewReader(os.Stdin)).Run()
}
