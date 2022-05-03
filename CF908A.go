package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF908A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF908A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF908A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF908A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF908A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF908A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF908A
Date: 1/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/908/A
Problem: CF908A
**/
func (in *CF908A) Run() {
	str := in.NextString()
	ans := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' ||
			str[i] == 'u' || str[i] == '1' || str[i] == '3' || str[i] == '5' || str[i] == '7' || str[i] == '9' {
			ans ++
		}
	}
	fmt.Println(ans)
}

func NewCF908A(r *bufio.Reader) *CF908A {
	return &CF908A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF908A(bufio.NewReader(os.Stdin)).Run()
}
