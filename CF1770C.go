package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1770C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1770C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1770C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1770C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1770C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1770C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1770C
Date: 12/30/2022
User: wotan
URL: https://codeforces.com/contests/1770/C
Problem: CF1770C
**/
func (in *CF1770C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt64()
		}

		x := int64(1)
		ans := true
		for x < 8601 {
			for i := 1; i < n && ans; i++ {
				for j := 0; j < i && ans; j++ {
					//fmt.Println(i, j, x, gcd(arr[i]+x, arr[j]+x), arr[i], arr[j])
					if gcd(arr[i]+x, arr[j]+x) != 1 {
						ans = false
					}
				}

			}
			if ans {
				break
			}

			x++
			if x < 8601 {
				ans = true
			}

		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func NewCF1770C(r *bufio.Reader) *CF1770C {
	return &CF1770C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1770C(bufio.NewReader(os.Stdin)).Run()
}
