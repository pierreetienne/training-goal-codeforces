package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1647A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1647A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1647A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1647A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1647A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1647A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1647A
Date: 22/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1647/A
Problem: CF1647A
**/
func (in *CF1647A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 1 {
			fmt.Println(1)
		} else if n == 2 {
			fmt.Println(2)
		} else {
			ans := ""
			if (n-1)%3 == 0 {
				sum := 0
				for i := 0; sum+3 <= n; i++ {
					ans = ans + "12"
					sum += 3
				}

				if sum < n {
					ans = ans + fmt.Sprintf("%d", n-sum)
				}

			} else {
				sum := 0
				for i := 0; sum+3 <= n; i++ {
					ans = ans + "21"
					sum += 3
				}

				if sum < n {
					ans = ans + fmt.Sprintf("%d", n-sum)
				}
			}
			fmt.Println(ans)
		}
	}
}

func NewCF1647A(r *bufio.Reader) *CF1647A {
	return &CF1647A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1647A(bufio.NewReader(os.Stdin)).Run()
}
