package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1671A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1671A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1671A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1671A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1671A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1671A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1671A
Date: 26/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1671/A
Problem: CF1671A
**/
func (in *CF1671A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		letter := uint8(0)
		count := 0
		ans := true
		for i := 0; i < len(str); i++ {
			if letter == str[i] {
				count++
			} else {
				if count != 0 && count < 2 {
					ans = false
					break
				}
				letter = str[i]
				count = 1
			}
		}
		if count != 0 && count < 2 {
			ans = false
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewCF1671A(r *bufio.Reader) *CF1671A {
	return &CF1671A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1671A(bufio.NewReader(os.Stdin)).Run()
}
