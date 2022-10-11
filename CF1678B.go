package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1678B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1678B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1678B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1678B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1678B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1678B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1678B
Date: 9/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1678/B1
Problem: CF1678B
**/
func (in *CF1678B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		arr := make([]int, 0)
		count := 0
		current := str[0]
		for i := 0; i < n; i++ {
			if str[i] == current {
				count++
			} else {
				arr = append(arr, count)
				count = 1
				current = str[i]
			}
		}
		arr = append(arr, count)

		ans := 0
		for i := 0; i < len(arr); i++ {
			if arr[i]%2 != 0 {
				arr[i]++
				ans++
				if i+1 < len(arr) {
					arr[i+1]--
					ans++
				}
			}
		}

		fmt.Println(ans / 2)

	}
}

func NewCF1678B(r *bufio.Reader) *CF1678B {
	return &CF1678B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1678B(bufio.NewReader(os.Stdin)).Run()
}
