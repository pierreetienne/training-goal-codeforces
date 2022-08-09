package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1716B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1716B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1716B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1716B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1716B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1716B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1716B
Date: 8/7/2022
User: pepradere
URL: https://codeforces.com/problemset/problem/1716/B
Problem: CF1716B
**/
func (in *CF1716B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		var ans strings.Builder
		pos := n
		ans.WriteString(fmt.Sprintf("%d\n", n))
		for j := 0; j < n; j++ {
			val := 1
			for i := 1; i <= n; i++ {
				if i == pos {
					ans.WriteString(fmt.Sprintf("%d ", n))
					pos--
				} else {
					ans.WriteString(fmt.Sprintf("%d ", val))
					val++
				}
			}
			ans.WriteString("\n")
		}
		fmt.Print(ans.String())
	}
}

func NewCF1716B(r *bufio.Reader) *CF1716B {
	return &CF1716B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1716B(bufio.NewReader(os.Stdin)).Run()
}
