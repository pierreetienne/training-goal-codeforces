package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1763A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1763A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1763A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1763A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1763A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1763A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1763A
Date: 28/12/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1763/A
Problem: CF1763A
**/
func (in *CF1763A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)

		max := 0
		indexMax := 0
		min := 1000000
		indexMin := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i] > max {
				max = arr[i]
				indexMax = i
			}
			if arr[i] < min {
				min = arr[i]
				indexMin = i
			}
		}

		for j := 0; j < 32; j++ {
			if max&(1<<j) == 0 {
				for i := 0; i < n; i++ {
					if i != indexMax && arr[i]&(1<<j) != 0 {
						arr[i] = arr[i] & (^(1 << j))
						max = max | (1 << j)
						break
					}
				}
			}
		}

		for j := 0; j < 32; j++ {
			if min&(1<<j) != 0 {
				for i := 0; i < n; i++ {
					if i != indexMin && arr[i]&(1<<j) == 0 {
						arr[i] = arr[i] | (1 << j)
						min = min & (^(1 << j))
						break
					}
				}
			}
		}

		fmt.Println(max - min)
	}
}

func NewCF1763A(r *bufio.Reader) *CF1763A {
	return &CF1763A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1763A(bufio.NewReader(os.Stdin)).Run()
}
