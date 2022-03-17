package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1430B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1430B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1430B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1430B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1430B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1430B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1430B
Date: 15/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1430/B
Problem: CF1430B
**/
func (in *CF1430B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt64()
		}
		sort.Slice(arr, func(i, j int) bool {
			if arr[i] < arr[j] {
				return true
			}
			return false
		})

		for i, j := n-2, 0; i >= 0 && j < k; i-- {
			arr[n-1] += arr[i]
			arr[i] = 0
			j++
		}
		sort.Slice(arr, func(i, j int) bool {
			if arr[i] < arr[j] {
				return true
			}
			return false
		})
		diff := int64(0)
		if n > 1 {
			diff = arr[n-1] - arr[0]
		}
		fmt.Println(diff)
	}
}

func NewCF1430B(r *bufio.Reader) *CF1430B {
	return &CF1430B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1430B(bufio.NewReader(os.Stdin)).Run()
}
