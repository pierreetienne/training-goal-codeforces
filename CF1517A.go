package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1517A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1517A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1517A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1517A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1517A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1517A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1517A
Date: 6/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1517/A
Problem: CF1517A
**/
func (in *CF1517A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		num := in.NextInt64()
		if num%2050 == 0 {
			val := fmt.Sprint(num/2050)
			sum := 0
			for i := 0; i < len(val); i++ {
				sum += int(val[i]-'0')
			}
			fmt.Println(sum)
		} else {
			fmt.Println(-1)
		}
	}
}

func NewCF1517A(r *bufio.Reader) *CF1517A {
	return &CF1517A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1517A(bufio.NewReader(os.Stdin)).Run()
}
