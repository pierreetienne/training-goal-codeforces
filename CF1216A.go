package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1216A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1216A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1216A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1216A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1216A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1216A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1216A
Date: 4/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1216/A
Problem: CF1216A
**/
func (in *CF1216A) Run() {
	n := in.NextInt()
	str := in.NextString()
	ans := 0
	var sol strings.Builder
	for i := 0; i < n; i += 2 {
		if str[i] == str[i+1] {
			sol.WriteString(string(str[i]))
			if str[i] == 'a' {
				sol.WriteString("b")
			} else {
				sol.WriteString("a")
			}
			ans++
		}else {
			sol.WriteString(string(str[i]))
			sol.WriteString(string(str[i+1]))
		}
	}
	fmt.Println(ans)
	fmt.Println(sol.String())
}

func NewCF1216A(r *bufio.Reader) *CF1216A {
	return &CF1216A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1216A(bufio.NewReader(os.Stdin)).Run()
}
