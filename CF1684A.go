package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1684A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1684A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1684A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1684A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1684A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1684A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1684A
Date: 21/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1684/A
Problem: CF1684A
**/
func (in *CF1684A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ans := -1
		if len(str) == 2 {
			ans = int(str[len(str)-1]-'0')
		}else {
			for i:=0;i<len(str);i++{
				if ans == -1 {
					ans = int(str[i]-'0')
				}else if ans >  int(str[i]-'0'){
					ans = int(str[i]-'0')
				}
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1684A(r *bufio.Reader) *CF1684A {
	return &CF1684A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1684A(bufio.NewReader(os.Stdin)).Run()
}
