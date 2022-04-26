package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1669C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1669C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1669C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1669C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1669C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1669C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1669C
Date: 24/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1669/C
Problem: CF1669C
**/
func (in *CF1669C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		oddodd := false
		oddeven := false

		eveneven := false
		evenodd := false
		ans := true
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if i%2 == 0 {
				if val%2 == 0 {
					oddodd = true
				} else {
					oddeven = true
				}
			} else {
				if val%2 == 0 {
					evenodd = true
				} else {
					eveneven = true
				}
			}
		}

		if oddeven && oddodd {
			ans = false
		}
		if eveneven && evenodd {
			ans = false
		}
		if ans{
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}

}

func NewCF1669C(r *bufio.Reader) *CF1669C {
	return &CF1669C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1669C(bufio.NewReader(os.Stdin)).Run()
}
