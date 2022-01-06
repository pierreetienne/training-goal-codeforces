package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1323A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1323A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1323A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1323A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1323A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1323A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1323A
Date: 2/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1323/A
Problem: CF1323A
**/
func (in *CF1323A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		pares := make([]int, 0)
		impares := make([]int, 0)
		for i := 0; i < n; i++ {
			m := in.NextInt()
			if m%2 == 0 {
				pares = append(pares, i+1)
			} else {
				impares = append(impares, i+1)
			}
		}
		ans := -1
		sol := make([]int, 0)
		if len(pares) > 0 {
			ans = 1
			sol = append(sol, pares[0])
		} else if len(impares) >= 2 {
			ans = 2
			sol = append(sol, impares[0])
			sol = append(sol, impares[1])
		}

		fmt.Println(ans)
		if ans != -1 {
			for i := 0; i < len(sol); i++ {
				if i == 0 {
					fmt.Print(sol[i])
				} else {
					fmt.Print(" ", sol[i])
				}
			}
		}
		fmt.Println()

	}
}

func NewCF1323A(r *bufio.Reader) *CF1323A {
	return &CF1323A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1323A(bufio.NewReader(os.Stdin)).Run()
}
