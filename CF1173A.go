package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1173A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1173A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1173A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1173A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1173A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1173A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1173A
Date: 22/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1173/A
Problem: CF1173A
**/
func (in *CF1173A) Run() {
	x, y, z := in.NextInt(), in.NextInt(), in.NextInt()

	if x+z > y && y+z < x {
		fmt.Println("+")
	} else if y+z > x && x+z < y {
		fmt.Println("-")
	} else if x+z == y+z && x+z == y && y+z == x {
		fmt.Println("0")
	} else {
		fmt.Println("?")
	}
}

func NewCF1173A(r *bufio.Reader) *CF1173A {
	return &CF1173A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1173A(bufio.NewReader(os.Stdin)).Run()
}
