package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF710A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF710A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF710A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF710A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF710A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF710A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF710A
Date: 17/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/710/A
Problem: CF710A
**/
func (in *CF710A) Run() {
	str := in.NextString()
	x := int(str[0] - 'a')
	y, _ := strconv.Atoi(string(str[1]))
	y -= 1

	if x > 0 && x < 7 && y > 0 && y < 7 {
		fmt.Println(8)
	} else if (x == 0 && y == 0) || (x == 0 && y == 7) || (x == 7 && y == 0) || (x == 7 && y == 7) {
		fmt.Println(3)
	} else {
		fmt.Println(5)
	}
}

func NewCF710A(r *bufio.Reader) *CF710A {
	return &CF710A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF710A(bufio.NewReader(os.Stdin)).Run()
}
