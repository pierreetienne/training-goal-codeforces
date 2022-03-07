package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1646A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1646A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1646A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1646A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1646A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1646A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1646A
Date: 5/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1646/A
Problem: CF1646A
**/
func (in *CF1646A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, s := in.NextInt64(), in.NextInt64()
		dn := n * n
		if s == 0 {
			fmt.Println(0)
		}else {
			sol := int64(math.Floor(float64(s / dn)))
			fmt.Println(sol)
		}


	}
}

func NewCF1646A(r *bufio.Reader) *CF1646A {
	return &CF1646A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1646A(bufio.NewReader(os.Stdin)).Run()
}
