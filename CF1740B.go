package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1740B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1740B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1740B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1740B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1740B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1740B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots
Run solve the problem CF1740B
Date: 11/17/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1740/B
Problem: CF1740B
**/
func (in *CF1740B) Run() {
	var str strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		sum := 0
		maxH := 0
		for i := 0; i < n; i++ {
			a, b := in.NextInt(), in.NextInt()
			if a < b {
				sum += 2 * a
				if maxH < b {
					maxH = b
				}
			} else {
				sum += 2 * b
				if maxH < a {
					maxH = a
				}
			}
		}
		sum += 2 * maxH

		str.WriteString(fmt.Sprintf("%v\n", sum))
	}
	fmt.Print(str.String())
}

func NewCF1740B(r *bufio.Reader) *CF1740B {
	return &CF1740B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1740B(bufio.NewReader(os.Stdin)).Run()
}
