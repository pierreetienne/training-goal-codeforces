package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF965A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF965A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF965A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF965A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF965A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF965A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF965A
Date: 12/25/2022
User: wotan
URL: https://codeforces.com/problemset/problem/965/A
Problem: CF965A
**/
func (in *CF965A) Run() {
	k, n, s, p := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
	fmt.Println(int(math.Ceil(float64(int(math.Ceil(float64(n)/float64(s)))*k) / float64(p))))

}

func NewCF965A(r *bufio.Reader) *CF965A {
	return &CF965A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF965A(bufio.NewReader(os.Stdin)).Run()
}
