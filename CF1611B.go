package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1611B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1611B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1611B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1611B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1611B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1611B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1611B
Date: 26/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1611/B
Problem: CF1611B
**/
func (in *CF1611B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()
		teams := int(math.Floor(float64(a+b) / 4.0))
		ans := 0
		if a >= teams && b >= teams {
			ans = teams
		} else {
			ans = int(math.Min(float64(a), float64(b)))
		}
		fmt.Println(ans)
	}
}

func NewCF1611B(r *bufio.Reader) *CF1611B {
	return &CF1611B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1611B(bufio.NewReader(os.Stdin)).Run()
}
