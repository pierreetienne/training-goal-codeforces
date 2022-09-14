package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1728B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1728B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1728B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1728B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1728B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1728B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1728B
Date: 9/10/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1728/B
Problem: CF1728B
**/
func (in *CF1728B) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 4 {
			fmt.Print("2 1 3 4")
		} else if n == 5 {
			arr := make([]int, n)
			arr[n-1] = n
			arr[n-2] = n - 1
			arr[n-3] = 3
			arr[n-4] = 2
			arr[n-5] = 1
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
		} else {
			arr := make([]int, n)
			arr[n-1] = n
			arr[n-2] = n - 1
			arr[n-3] = 1
			arr[n-4] = 3
			arr[n-5] = 2
			for i, j := 4, 0; i < n-1; i++ {
				arr[j] = i
				j++
			}
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
		}
		fmt.Println()
	}

}

func NewCF1728B(r *bufio.Reader) *CF1728B {
	return &CF1728B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1728B(bufio.NewReader(os.Stdin)).Run()
}
