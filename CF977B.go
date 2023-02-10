package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF977B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF977B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF977B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF977B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF977B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF977B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots team sdk
Run solve the problem CF977B
Date: 2/3/2023
User: wotan
URL: https://codeforces.com/problemset/problem/977/B
Problem: CF977B
**/
func (in *CF977B) Run() {
	n := in.NextInt()
	s := in.NextString()

	res := 0
	ans := ""
	for i := 0; i < n-1; i++ {
		cur := 0
		for j := 0; j < n-1; j++ {
			if s[j] == s[i] && s[j+1] == s[i+1] {
				cur++
			}
			if res < cur {
				res = cur
				ans = string(s[i]) + string(s[i+1])
			}
		}
	}
	fmt.Println(ans)
}

func NewCF977B(r *bufio.Reader) *CF977B {
	return &CF977B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF977B(bufio.NewReader(os.Stdin)).Run()
}
