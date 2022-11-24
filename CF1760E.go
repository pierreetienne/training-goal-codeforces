package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1760E struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1760E) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1760E) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1760E) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1760E) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1760E) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1760E
Date: 11/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1760/E
Problem: CF1760E
**/
func (in *CF1760E) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)

		leftZero := -1
		rightOne := -1
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i] == 0 && leftZero == -1 {
				leftZero = i
			}

			if arr[i] == 1 {
				rightOne = i
			}
		}
		c := fu(&arr, n)

		a := int64(0)
		if leftZero != -1 {
			arr[leftZero] = 1
			a = fu(&arr, n)
			arr[leftZero] = 0
		}

		b := int64(0)
		if rightOne != -1 {
			arr[rightOne] = 0
			b = fu(&arr, n)
		}

		fmt.Println(int64(math.Max(float64(a), math.Max(float64(b), float64(c)))))

	}
}

func fu(arr *[]int, n int) int64 {
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			dp[i] = dp[i+1]
		}
		if (*arr)[i] == 0 {
			dp[i]++
		}

	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		if (*arr)[i] == 1 {
			ans += int64(dp[i])
		}
	}
	return ans
}

func NewCF1760E(r *bufio.Reader) *CF1760E {
	return &CF1760E{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1760E(bufio.NewReader(os.Stdin)).Run()
}
