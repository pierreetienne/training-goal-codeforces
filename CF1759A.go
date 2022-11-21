package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1759A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1759A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1759A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1759A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1759A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1759A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1759A
Date: 11/19/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1759/A
Problem: CF1759A
**/
func (in *CF1759A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		if str[0] == 'e' {
			str = "Y" + str
		} else if str[0] == 's' {
			str = "Ye" + str
		}

		if str[len(str)-1] == 'e' {
			str = str + "s"
		} else if str[len(str)-1] == 'Y' {
			str = str + "es"
		}

		ans := true
		for i := 2; i < len(str); i += 3 {
			if str[i-2] != 'Y' || str[i-1] != 'e' || str[i] != 's' {
				ans = false
			}
		}
		if ans && len(str)%3 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1759A(r *bufio.Reader) *CF1759A {
	return &CF1759A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1759A(bufio.NewReader(os.Stdin)).Run()
}
