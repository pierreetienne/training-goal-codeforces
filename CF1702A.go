package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1702A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1702A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1702A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1702A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1702A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1702A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1702A
Date: 12/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1702/A
Problem: CF1702A
**/
func (in *CF1702A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		m := in.NextString()
		v, _ := strconv.ParseInt(m, 10, 64)
		b := int64(math.Pow(float64(10), float64(len(m)-1)))
		fmt.Println(v-b)
	}
}

func NewCF1702A(r *bufio.Reader) *CF1702A {
	return &CF1702A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1702A(bufio.NewReader(os.Stdin)).Run()
}
