package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1561A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1561A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1561A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1561A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1561A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1561A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1561A
Date: 11/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1561/A
Problem: CF1561A
**/
func (in *CF1561A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		sol := 0
		arr := make([]int, n)

		order := true
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i] != i+1 {
				order = false
			}
		}

		ite := 0

		for {
			ws := true
			for i:= 0;i<n;i++{
				if arr[i] != i+1{
					ws = false
				}
			}
			for i := ite; i < n-1; i += 2 {

				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]

				}
			}
			if ite == 0 {
				ite = 1
			} else {
				ite = 0
			}

			sol++
			if ws{
				break
			}


		}

		if order {
			fmt.Println(0)
		} else {
			fmt.Println(sol-1)
		}

	}
}

func NewCF1561A(r *bufio.Reader) *CF1561A {
	return &CF1561A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1561A(bufio.NewReader(os.Stdin)).Run()
}
