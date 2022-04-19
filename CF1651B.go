package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1651B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1651B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1651B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1651B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1651B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1651B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1651B
Date: 18/04/22
User: epradere
URL: https://codeforces.com/problemset/problem/1651/B
Problem: CF1651B
**/
func (in *CF1651B) Run() {
	var str strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n > 19 {
			str.WriteString("NO\n")
		} else {
			str.WriteString("YES\n")
			value := int64(1)
			for i := 0; i < n; i++ {
				str.WriteString(fmt.Sprintf("%v ", value))
				value *= 3
			}
			str.WriteString("\n")
		}
	}
	fmt.Println(str.String())
}

func NewCF1651B(r *bufio.Reader) *CF1651B {
	return &CF1651B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1651B(bufio.NewReader(os.Stdin)).Run()
}