package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1646B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1646B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1646B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1646B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1646B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1646B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1646B
Date: 2/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1646/B
Problem: CF1646B
**/
func (in *CF1646B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		ans := false
		a := int64(arr[0] + arr[1])
		b := int64(arr[n-1])
		i, j := 2, n-2

		if a < b {
			ans = true
		}

		for a >= b && i < j && !ans {

			a += int64(arr[i])
			b += int64(arr[j])
			j--
			i++
			if a < b {
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

func NewCF1646B(r *bufio.Reader) *CF1646B {
	return &CF1646B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1646B(bufio.NewReader(os.Stdin)).Run()
}
