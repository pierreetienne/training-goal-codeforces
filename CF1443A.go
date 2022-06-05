package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1443A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1443A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1443A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1443A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1443A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1443A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1443A
Date: 3/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1443/A
Problem: CF1443A
**/
func (in *CF1443A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		num := n*2
		for i:=0;i<n;i++{
			fmt.Print(num, " ")
			num+=2
		}
		fmt.Println()
	}
}

func NewCF1443A(r *bufio.Reader) *CF1443A {
	return &CF1443A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1443A(bufio.NewReader(os.Stdin)).Run()
}
