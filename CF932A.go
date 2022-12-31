package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF932A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF932A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF932A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF932A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF932A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF932A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Gopherbots NEWWWWW ez manejo slices
Run solve the problem CF932A
Date: 11/29/2022
User: wotan
URL: https://codeforces.com/problemset/problem/932/A
Problem: CF932A
**/
func (in *CF932A) Run() {
	str := in.NextString()

	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			str = str[0:j+1] + string(str[i]) + str[j+1:]
			j++
		} else {
			j--
			i++
		}
	}
	fmt.Println(str)
}

func NewCF932A(r *bufio.Reader) *CF932A {
	return &CF932A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF932A(bufio.NewReader(os.Stdin)).Run()
}
