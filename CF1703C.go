package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1703C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1703C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1703C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1703C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1703C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1703C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
solved 46ms gopherbots
Run solve the problem CF1703C
Date: 14/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1703/C
Problem: CF1703C
**/
func (in *CF1703C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		w := make([]int, n)
		for i := 0; i < n; i++ {
			w[i] = in.NextInt()
		}

		ans := make([]int, n)
		for i := 0; i < n; i++ {
			in.NextString()
			str := in.NextString()
			ans[i] = w[i]
			for j := 0; j < len(str); j++ {
				if str[j] == 'D' {
					ans[i]= (ans[i]+1)%10
				}else {
					ans[i] = ((10+ans[i])-1)%10
				}
			}
		}

		for i:=0;i<n;i++{
			fmt.Print(ans[i], " ")
		}
		fmt.Println()
	}
}

func NewCF1703C(r *bufio.Reader) *CF1703C {
	return &CF1703C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1703C(bufio.NewReader(os.Stdin)).Run()
}
