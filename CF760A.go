package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF760A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF760A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF760A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF760A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF760A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF760A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF760A
Date: 11/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/760/A
Problem: CF760A
**/
func (in *CF760A) Run() {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	m, d := in.NextInt(), in.NextInt()-1

	val := int(math.Ceil(float64(months[m-1]+d) / 7.0))

	fmt.Println(val)

}

func NewCF760A(r *bufio.Reader) *CF760A {
	return &CF760A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF760A(bufio.NewReader(os.Stdin)).Run()
}
