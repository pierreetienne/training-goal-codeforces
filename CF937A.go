package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF937A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF937A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF937A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF937A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF937A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF937A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF937A
Date: 23/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/937/A
Problem: CF937A
**/
func (in *CF937A) Run() {
	n := in.NextInt()
	mapa := map[int]int{}

	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val > 0 {
			mapa[val] = 1
		}
	}
	fmt.Println(len(mapa))

}

func NewCF937A(r *bufio.Reader) *CF937A {
	return &CF937A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF937A(bufio.NewReader(os.Stdin)).Run()
}
