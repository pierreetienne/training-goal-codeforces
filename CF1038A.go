package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1038A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1038A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1038A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1038A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1038A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1038A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1038A
Date: 14/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1038/A
Problem: CF1038A
**/
func (in *CF1038A) Run() {
	n, k := in.NextInt(), in.NextInt()
	str := in.NextString()
	mapa := make(map[uint8]int)
	for i := 0; i < n; i++ {
		mapa[str[i]]++
	}

	min := n

	for i := 0; i < k; i++ {
		num := uint8('A' + i)
		if min > mapa[num] {
			min = mapa[num]
		}
	}

	fmt.Println(min * k)

}

func NewCF1038A(r *bufio.Reader) *CF1038A {
	return &CF1038A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1038A(bufio.NewReader(os.Stdin)).Run()
}
