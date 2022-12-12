package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1130A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1130A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1130A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1130A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1130A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1130A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1130A
Date: 12/9/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1130/A
Problem: CF1130A
**/
func (in *CF1130A) Run() {

	n := in.NextInt()
	num := 0
	neg := 0

	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val > 0 {
			num++
		}

		if val < 0 {
			neg++
		}
	}

	d := int(math.Ceil(float64(n) / 2.))

	if num >= d {
		fmt.Println(1)
	} else if neg >= d {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}

}

func NewCF1130A(r *bufio.Reader) *CF1130A {
	return &CF1130A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1130A(bufio.NewReader(os.Stdin)).Run()
}
