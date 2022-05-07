package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1635B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1635B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1635B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1635B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1635B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1635B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1635B
Date: 5/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1635/B
Problem: CF1635B
**/
func (in *CF1635B) Run() {
	var sol strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		ans := 0
		for i := 1; i < n-1; i++ {
			if arr[i] > arr[i+1] && arr[i] > arr[i-1] {
				if i+2 < n {
					arr[i+1] = int(math.Max(float64(arr[i]), float64(arr[i+2])))
				} else {
					arr[i+1] = arr[i]
				}
				ans++
			}
		}

		sol.WriteString(fmt.Sprintf("%d\n", ans))
		for i := 0; i < n; i++ {
			sol.WriteString(fmt.Sprintf("%d ", arr[i]))
		}
		sol.WriteString("\n")
	}
	fmt.Println(sol.String())
}

func NewCF1635B(r *bufio.Reader) *CF1635B {
	return &CF1635B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1635B(bufio.NewReader(os.Stdin)).Run()
}
