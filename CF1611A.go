package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1611A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1611A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1611A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1611A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1611A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1611A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1611A
Date: 1/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1611/A
Problem: CF1611A
**/
func (in *CF1611A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		num := in.NextInt()
		ans := 0
		if num%2 != 0 {
			ans = -1
			str := strconv.Itoa(num)
			if (str[0]-'0')%2 == 0 {
				ans = 1
			} else {
				for i := 0; i < len(str); i++ {
					if (str[i]-'0')%2 == 0 {
						ans = 2
					}
				}
			}
		}
		fmt.Println(ans)

	}
}

func NewCF1611A(r *bufio.Reader) *CF1611A {
	return &CF1611A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1611A(bufio.NewReader(os.Stdin)).Run()
}
