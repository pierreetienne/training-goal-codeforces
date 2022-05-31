package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1682A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1682A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1682A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1682A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1682A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1682A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1682A
Date: 29/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1682/A
Problem: CF1682A
**/
func (in *CF1682A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		mid := int(math.Floor(float64(n) / 2.))

		left := mid
		for i := mid; i >= 0; i-- {
			if str[mid] == str[i] {
				left = i
			} else {
				break
			}
		}

		right := mid
		for i := mid; i < n; i++ {
			if str[mid] == str[i] {
				right = i
			} else {
				break
			}
		}
		fmt.Println((right - left) + 1)

	}
}

func NewCF1682A(r *bufio.Reader) *CF1682A {
	return &CF1682A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1682A(bufio.NewReader(os.Stdin)).Run()
}
