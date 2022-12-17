package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF669A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF669A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF669A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF669A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF669A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF669A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF669A
Date: 12/14/2022
User: wotan
URL: https://codeforces.com/problemset/problem/669/A
Problem: CF669A
**/
func (in *CF669A) Run() {
	n := in.NextInt()
	if n == 1 {
		fmt.Println(1)
	} else if n == 2 {
		fmt.Println(1)
	} else {
		f := int(math.Floor(float64(n) / 3.))
		ans := 1 + 2*f
		if n%3 == 0 {
			ans--
		}
		fmt.Println(ans)
	}

}

func NewCF669A(r *bufio.Reader) *CF669A {
	return &CF669A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF669A(bufio.NewReader(os.Stdin)).Run()
}
