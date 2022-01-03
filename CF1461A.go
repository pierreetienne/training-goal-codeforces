package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1461A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1461A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1461A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1461A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1461A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1461A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1461A
Date: 2/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1461/A
Problem: CF1461A
**/
func (in *CF1461A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		letters := []string{"a", "b", "c"}
		var sb strings.Builder
		for i := 0;i < k;i++{
			sb.WriteString(letters[0])
		}
		for i := 0; i < n-k; i++ {
			sb.WriteString(letters[(i+1)%3])
		}
		fmt.Println(sb.String())
	}
}

func NewCF1461A(r *bufio.Reader) *CF1461A {
	return &CF1461A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1461A(bufio.NewReader(os.Stdin)).Run()
}
