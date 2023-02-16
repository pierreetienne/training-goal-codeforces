package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1709B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1709B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1709B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1709B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1709B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1709B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1709B
Date: 12/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1709/B
Problem: CF1709B
**/
func (in *CF1709B) Run() {
	n, m := in.NextInt(), in.NextInt()
	arr := make([]int, n)
	dp := make([]int, n)
	dpL := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
		if i > 0 {
			if arr[i-1] > arr[i] {
				dp[i] = arr[i-1] - arr[i]
			}
			dp[i] += dp[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i+1] > arr[i] {
			dpL[i] = arr[i+1] - arr[i]
		}
		dpL[i] += dpL[i+1]
	}

	var sol strings.Builder
	for i := 0; i < m; i++ {
		s, t := in.NextInt(), in.NextInt()
		if s < t {
			sol.WriteString(fmt.Sprintf("%d\n", dp[t-1]-dp[s-1]))
		} else {
			sol.WriteString(fmt.Sprintf("%d\n", dpL[t-1]-dpL[s-1]))
		}
	}
	fmt.Print(sol.String())
}

func NewCF1709B(r *bufio.Reader) *CF1709B {
	return &CF1709B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1709B(bufio.NewReader(os.Stdin)).Run()
}
