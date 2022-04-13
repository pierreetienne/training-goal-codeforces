package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1656A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1656A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1656A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1656A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1656A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1656A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1656A
Date: 11/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1656/A
Problem: CF1656A
**/
func (in *CF1656A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		min , max := 0, 0
		indexmin , indexmax := 0, 0
		for i:=0;i<n;i++{
			val := in.NextInt()
			if i == 0{
				min = val
				max = val
				indexmax = i+1
				indexmin = i+1
			}

			if val < min {
				min = val
				indexmin = i+1
			}

			if val > max {
				max = val
				indexmax = i+1
			}
		}
		fmt.Println(indexmin,indexmax)
	}
}

func NewCF1656A(r *bufio.Reader) *CF1656A {
	return &CF1656A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1656A(bufio.NewReader(os.Stdin)).Run()
}
