package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1462B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1462B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1462B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1462B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1462B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1462B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1462B
Date: 27/10/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1462/B
Problem: CF1462B
**/
func (in *CF1462B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		ws := false
		for i := 0; i < n && !ws; i++ {
			for j := i; j < n && !ws; j++ {
				nstr := str[0:i]+str[j:n]
				fmt.Println(i, j , nstr)
				if nstr == "2020" {
					ws = true
				}
			}
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1462B(r *bufio.Reader) *CF1462B {
	return &CF1462B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1462B(bufio.NewReader(os.Stdin)).Run()
}
