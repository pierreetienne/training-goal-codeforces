package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1719B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1719B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1719B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1719B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1719B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1719B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1719B
Date: 9/14/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1719/B
Problem: CF1719B
**/
func (in *CF1719B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt64(), in.NextInt64()
		if k%2 == 0 {
			c := 0
			var ans strings.Builder
			ans.WriteString("YES\n")
			ws := true
			for i := int64(0); i < n; i += 2 {
				if c%2 == 0 {
					if !f(i+2, i+1, k) {
						ws = false
						break
					}
					ans.WriteString(fmt.Sprintf("%v %v\n", i+2, i+1))

				} else {
					if !f(i+1, i+2, k) {
						ws = false
						break
					}
					ans.WriteString(fmt.Sprintf("%v %v\n", i+1, i+2))
				}
				c++
			}
			if ws {
				fmt.Print(ans.String())
			} else {
				fmt.Println("NO")
			}
		} else if k%2 != 0 {
			var ans strings.Builder
			ans.WriteString("YES\n")
			ws := true
			for i := int64(0); i < n; i += 2 {
				if !f(i+1, i+2, k) {
					ws = false
					break
				}
				ans.WriteString(fmt.Sprintf("%v %v\n", i+1, i+2))
			}
			if ws {
				fmt.Print(ans.String())
			} else {
				fmt.Println("NO")
			}
		}

	}
}

func f(a, b, k int64) bool {
	return ((a+k)*b)%4 == 0
}

func NewCF1719B(r *bufio.Reader) *CF1719B {
	return &CF1719B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1719B(bufio.NewReader(os.Stdin)).Run()
}
