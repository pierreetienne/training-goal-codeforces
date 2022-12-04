package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1143A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1143A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1143A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1143A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1143A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1143A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1143A
Date: 12/2/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1143/A
Problem: CF1143A
**/
func (in *CF1143A) Run() {
	n := in.NextInt()
	lastZero := 0
	lastOne := 0
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val == 1 {
			lastOne = i
		} else {
			lastZero = i
		}
	}
	fmt.Println(int(math.Min(float64(lastZero+1), float64(lastOne+1))))
}

func NewCF1143A(r *bufio.Reader) *CF1143A {
	return &CF1143A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1143A(bufio.NewReader(os.Stdin)).Run()
}
