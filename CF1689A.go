package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1689A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1689A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1689A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1689A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1689A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1689A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1689A
Date: 2/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1689/A
Problem: CF1689A
**/
func (in *CF1689A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m, k := in.NextInt(), in.NextInt(), in.NextInt()
		g := in.NextString()
		h := in.NextString()

		a := make([]uint8, n)
		b := make([]uint8, m)

		for i := 0; i < n; i++ {
			a[i] = g[i]
		}

		sort.Slice(a, func(i, j int) bool {
			if a[i] < a[j] {
				return true
			}
			return false
		})

		for i := 0; i < m; i++ {
			b[i] = h[i]
		}

		sort.Slice(b, func(i, j int) bool {
			if b[i] < b[j] {
				return true
			}
			return false
		})

		sol := ""
		for i, j, p, q := 0, 0, 0, 0; ; {
			x := uint8('-')
			if i < n {
				x = a[i]
			}

			y := uint8('-')
			if j < m {
				y = b[j]
			}
			if x != '-' && y != '-' {
				if x < y {
					if p < k {
						sol += string(rune(x))
						p++
						q = 0
						i++
					} else {
						sol += string(rune(y))
						p = 0
						q = 1
						j++
					}
				} else {
					if q < k {
						sol += string(rune(y))
						q++
						j++
						p = 0
					} else {
						sol += string(rune(x))
						q = 0
						p = 1
						i++
					}
				}
			} else {
				break
			}
		}
		fmt.Println(sol)

	}
}

func NewCF1689A(r *bufio.Reader) *CF1689A {
	return &CF1689A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1689A(bufio.NewReader(os.Stdin)).Run()
}
