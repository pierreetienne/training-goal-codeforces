package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1747A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1747A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1747A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1747A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1747A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1747A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1747A
Date: 5/11/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1747/A
Problem: CF1747A
**/
func (in *CF1747A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		pos := 0
		neg := 0
		all := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if val > 0 {
				pos += val
			} else {
				neg += val
			}
			all += val
		}

		max := int(math.Max(float64(all), math.Min(-float64(neg)-float64(pos), -float64(pos)-float64(neg))))
		fmt.Println(max)
	}
}

func NewCF1747A(r *bufio.Reader) *CF1747A {
	return &CF1747A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1747A(bufio.NewReader(os.Stdin)).Run()
}
