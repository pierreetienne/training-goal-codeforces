package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1698B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1698B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1698B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1698B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1698B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1698B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1698B
Date: 4/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1698/B
Problem: CF1698B
**/
func (in *CF1698B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}

		ans := 0

		if k == 1 {
			ans = int(math.Ceil(float64(n-2) / 2))
		}else {
			for i := 1; i < n-1; i++ {
				if arr[i-1]+arr[i+1] < arr[i] {
					ans++
				}
			}
		}

		fmt.Println(ans)

	}
}

func NewCF1698B(r *bufio.Reader) *CF1698B {
	return &CF1698B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1698B(bufio.NewReader(os.Stdin)).Run()
}
