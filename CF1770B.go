package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1770B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1770B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1770B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1770B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1770B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1770B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1770B
Date: 12/30/2022
User: wotan
URL: https://codeforces.com/contests/1770/B
Problem: CF1770B
**/
func (in *CF1770B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, _ := in.NextInt(), in.NextInt()
		var ans strings.Builder
		for i, j := 1, n; i <= j; i++ {
			if i == j {
				ans.WriteString(fmt.Sprintf("%d", i))
			} else {
				ans.WriteString(fmt.Sprintf("%d %d ", j, i))
			}

			j--
		}
		fmt.Println(ans.String())
	}
}

func NewCF1770B(r *bufio.Reader) *CF1770B {
	return &CF1770B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1770B(bufio.NewReader(os.Stdin)).Run()
}
