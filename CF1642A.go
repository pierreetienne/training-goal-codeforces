package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1642A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1642A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1642A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1642A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1642A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1642A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1642A
Date: 9/17/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1642/A
Problem: CF1642A
**/
func (in *CF1642A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		m := make([][]float64, 3)
		m[0] = []float64{
			float64(in.NextInt()), float64(in.NextInt()),
		}
		m[1] = []float64{
			float64(in.NextInt()), float64(in.NextInt()),
		}
		m[2] = []float64{
			float64(in.NextInt()), float64(in.NextInt()),
		}
		ans := 0.0

		if m[0][1] == m[1][1] && m[0][1] > m[2][1] {
			ans = math.Abs(m[0][0] - m[1][0])
		} else if m[0][1] == m[2][1] && m[0][1] > m[1][1] {
			ans = math.Abs(m[0][0] - m[2][0])
		} else if m[1][1] == m[2][1] && m[1][1] > m[0][1] {
			ans = math.Abs(m[1][0] - m[2][0])
		}

		fmt.Printf("%f\n", ans)
	}
}

func NewCF1642A(r *bufio.Reader) *CF1642A {
	return &CF1642A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1642A(bufio.NewReader(os.Stdin)).Run()
}
