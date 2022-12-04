package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF15884A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF15884A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF15884A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF15884A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF15884A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF15884A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF15884A
Date: 8/11/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1584/A
Problem: CF15884A
**/
func (in *CF15884A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		u, v := in.NextInt64(), in.NextInt64()
		fmt.Println(-(u * u), v*v)
	}
}

func NewCF15884A(r *bufio.Reader) *CF15884A {
	return &CF15884A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF15884A(bufio.NewReader(os.Stdin)).Run()
}
