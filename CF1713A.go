package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1713A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1713A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1713A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1713A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1713A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1713A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1713A
Date: 8/13/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1713/A
Problem: CF1713A
**/
var arr [][]int

func (in *CF1713A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()

		minX := 0
		minY := 0
		maxX := 0
		maxY := 0

		for i := 0; i < n; i++ {
			x := in.NextInt()
			y := in.NextInt()
			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}

			if x < minX {
				minX = x
			}
			if y < minY {
				minY = y
			}
		}
		ans := 2 * (maxX + maxY - minX - minY)
		fmt.Println(ans)
	}
}

func f(x, y int, v []bool) int {
	min := -1
	for i := 0; i < len(v); i++ {
		if !v[i] {
			v[i] = true
			ans := f(arr[i][0], arr[i][1], v) + distance(x, y, arr[i][0], arr[i][1])
			if ans < min || min == -1 {
				min = ans
			}
			v[i] = false
		}
	}

	if min == -1 {
		return distance(x, y, 100, 100)
	}

	return min
}

func distance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2))) + int(math.Abs(float64(y1)-float64(y2)))
}

func NewCF1713A(r *bufio.Reader) *CF1713A {
	return &CF1713A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1713A(bufio.NewReader(os.Stdin)).Run()
}
