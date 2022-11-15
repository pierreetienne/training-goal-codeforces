package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1271A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1271A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1271A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1271A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1271A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1271A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1271A
Date: 1/11/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1271/A
Problem: CF1271A
**/
func (in *CF1271A) Run() {
	a, b, c, d, e, f := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
	x := int(math.Min(float64(a), float64(d)))
	y := int(math.Min(float64(b), math.Min(float64(c), float64(d))))

	val := (x * e) + (int(math.Min(float64(b), math.Min(float64(c), float64(d-x)))) * f)

	valB := (int(math.Min(float64(a), float64(d-y))) * e) + (y * f)

	fmt.Println(int(math.Max(float64(valB), float64(val))))
}

func NewCF1271A(r *bufio.Reader) *CF1271A {
	return &CF1271A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1271A(bufio.NewReader(os.Stdin)).Run()
}
