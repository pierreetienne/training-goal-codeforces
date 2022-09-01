package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1721A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1721A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1721A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1721A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1721A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1721A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1721A
Date: 8/29/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1721/A
Problem: CF1721A
**/
func (in *CF1721A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		mapa := make(map[uint8]bool)
		str := in.NextString()
		mapa[str[0]] = true
		mapa[str[1]] = true
		str = in.NextString()
		mapa[str[0]] = true
		mapa[str[1]] = true
		if len(mapa) == 1 {
			fmt.Println(0)
		} else if len(mapa) == 2 {
			fmt.Println(1)
		} else if len(mapa) == 3 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}

func NewCF1721A(r *bufio.Reader) *CF1721A {
	return &CF1721A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1721A(bufio.NewReader(os.Stdin)).Run()
}
