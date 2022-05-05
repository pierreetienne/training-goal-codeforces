package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1452C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1452C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1452C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1452C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1452C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1452C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1452C
Date: 9/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1452/C
Problem: CF1452C
**/
func (in *CF1452C) Run() {
	for n := in.NextInt(); n > 0; n-- {
		str := in.NextString()
		p := 0
		b := 0
		ans :=0
		for i := 0; i < len(str); i++ {
			if str[i] == '(' {
				p++
			}else if str[i] == '[' {
				b++
			}else if str[i] == ')' {
				if p > 0 {
					ans++
					p--
				}
			}else if str[i] == ']' {
				if b > 0 {
					ans++
					b--
				}
			}
		}

		fmt.Println(ans)

	}
}

func NewCF1452C(r *bufio.Reader) *CF1452C {
	return &CF1452C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1452C(bufio.NewReader(os.Stdin)).Run()
}
