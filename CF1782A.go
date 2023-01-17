package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1782A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1782A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1782A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1782A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1782A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1782A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1782A
Date: 1/15/2023
User: wotan
URL: https://codeforces.com/contest/1782/problem/A
Problem: CF1782A
**/
func (in *CF1782A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		w, d, h := in.NextInt(), in.NextInt(), in.NextInt()
		a, b, f, g := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
		arr := []int{
			b + h + g + int(math.Abs(float64(a)-float64(f))),
			int(math.Abs(float64(d)-float64(b))) + h + int(float64(d)-float64(g)) + int(math.Abs(float64(a)-float64(f))),
			a + h + f + int(math.Abs(float64(g)-float64(b))),
			int(math.Abs(float64(w)-float64(a))) + h + int(math.Abs(float64(w)-float64(f))) + int(math.Abs(float64(b)-float64(g))),
		}
		sort.Ints(arr)
		fmt.Println(arr[0])
	}
}

func NewCF1782A(r *bufio.Reader) *CF1782A {
	return &CF1782A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1782A(bufio.NewReader(os.Stdin)).Run()
}
