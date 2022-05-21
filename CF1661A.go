package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1661A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1661A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1661A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1661A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1661A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1661A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1661A
Date: 19/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1661/A
Problem: CF1661A
**/
func (in *CF1661A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}

		b := make([]int, n)
		for i := 0; i < n; i++ {
			b[i] = in.NextInt()
		}

		sum := int64(0)
		for i := 0; i < n; i++ {
			if a[i] > b[i] {
				a[i], b[i] = b[i], a[i]
			}

		}

		for i:=1;i<n;i++{
			sum += int64(math.Abs(float64(a[i-1])-float64(a[i])))
			sum += int64(math.Abs(float64(b[i-1])-float64(b[i])))
		}
		fmt.Println(sum)
	}
}

func NewCF1661A(r *bufio.Reader) *CF1661A {
	return &CF1661A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1661A(bufio.NewReader(os.Stdin)).Run()
}
