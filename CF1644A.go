package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1644A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1644A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1644A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1644A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1644A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1644A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1644A
Date: 6/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1644/A
Problem: CF1644A
**/
func (in *CF1644A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		r,g,b := false, false, false
		ans := true
		for i:=0;i<len(str);i++ {
			if str[i] == 'r' {
				r = true
			}else if str[i] == 'g' {
				g = true
			} else if str[i] == 'b' {
				b = true
			} else if str[i] == 'R' {
				if !r {
					ans = false
				}else {
					r = false
				}
			}else if str[i] == 'G' {
				if !g {
					ans = false
				}else {
					g = false
				}
			} else if str[i] == 'B' {
				if !b {
					ans = false
				} else {
					b = false
				}
			}
		}
		if ans {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}
}

func NewCF1644A(r *bufio.Reader) *CF1644A {
	return &CF1644A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1644A(bufio.NewReader(os.Stdin)).Run()
}
