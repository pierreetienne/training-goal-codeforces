package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1688A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1688A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1688A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1688A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1688A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1688A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1688A
Date: 14/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1688/A
Problem: CF1688A
**/
func (in *CF1688A) Run() {
	var str strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		ans := 0

		count := 0
		for i:=0;i<31;i++{
			if n&(1<<i) != 0 {
				count++
			}
		}

		and := false
		xor := count > 1

		for i := 0; i <= 30; i++ {
			if n&(1<<i) != 0 {
				if !and {
					ans = ans | (1 << i)
					and = true
				}
			}else {
				if !xor {
					ans = ans | (1 << i)
					xor = true
				}
			}


			if xor && and {
				break
			}
		}
		str.WriteString(fmt.Sprintf("%d\n", ans))
	}
	fmt.Print(str.String())
}

func NewCF1688A(r *bufio.Reader) *CF1688A {
	return &CF1688A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1688A(bufio.NewReader(os.Stdin)).Run()
}
