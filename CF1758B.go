package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1758B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1758B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1758B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1758B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1758B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1758B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1758B
Date: 2/10/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1758/B
Problem: CF1758B
**/
func (in *CF1758B) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		if n%2 != 0 {
			for i := 0; i < n; i++ {
				ans.WriteString(fmt.Sprint((n*2)+1, " "))
			}
		} else {
			ans.WriteString(fmt.Sprint(1, 3, " "))
			for i := 0; i < n-2; i++ {
				ans.WriteString(fmt.Sprint(2, " "))
			}
		}
		ans.WriteString("\n")
	}
	fmt.Print(ans.String())
}

func NewCF1758B(r *bufio.Reader) *CF1758B {
	return &CF1758B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1758B(bufio.NewReader(os.Stdin)).Run()
}
