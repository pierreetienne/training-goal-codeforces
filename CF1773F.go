package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1773F struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1773F) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1773F) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1773F) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1773F) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1773F) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1773F
Date: 2/1/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1773/F
Problem: CF1773F
**/
func (in *CF1773F) Run() {
	n := in.NextInt()
	a, b := in.NextInt(), in.NextInt()
	swap := false
	if a < b {
		swap = true
		a, b = b, a
	}

	sa := 0
	sb := 0
	arr := make([][]int, n)
	for i := 0; i < n-1; i++ {
		if sa < a {
			arr[i] = []int{1, 0}
			sa++
		} else if sa == a {
			if sb < b {
				arr[i] = []int{0, 1}
				sb++
			}
		}
	}

	var str strings.Builder
	if n-2 >= 0 && arr[n-2] != nil && a-sa == b-sb && a-sa > 0 {
		arr[n-2][0]++
		sa++
	}
	arr[n-1] = []int{a - sa, b - sb}
	tie := 0
	for i := 0; i < n; i++ {
		if arr[i] == nil {
			str.WriteString("0:0\n")
			tie++
		} else {
			if arr[i][0] == arr[i][1] {
				tie++

			}
			if !swap {
				str.WriteString(fmt.Sprintf("%d:%d\n", arr[i][0], arr[i][1]))
			} else {
				str.WriteString(fmt.Sprintf("%d:%d\n", arr[i][1], arr[i][0]))
			}

		}
	}

	fmt.Println(tie)
	fmt.Print(str.String())
}

func NewCF1773F(r *bufio.Reader) *CF1773F {
	return &CF1773F{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1773F(bufio.NewReader(os.Stdin)).Run()
}
