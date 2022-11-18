package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1747B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1747B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1747B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1747B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1747B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1747B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1747B
Date: 11/16/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1747/B
Problem: CF1747B
**/
func (in *CF1747B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 1 {
			fmt.Println("1\n1 2")
		} else {
			var sol strings.Builder
			count := 0
			for i := 0; i < n/2; i++ {
				count++
				x := (3 * i) + 1
				y := 3 * (n - i)

				sol.WriteString(fmt.Sprintf("%d %d\n", x, y))
			}

			if n%2 != 0 {
				count++
				x := ((3 * n) - 1) / 2
				y := x + 2
				sol.WriteString(fmt.Sprintf("%d %d\n", x, y))
			}

			fmt.Println(count)
			fmt.Print(sol.String())
		}
	}
}

func NewCF1747B(r *bufio.Reader) *CF1747B {
	return &CF1747B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1747B(bufio.NewReader(os.Stdin)).Run()
}
