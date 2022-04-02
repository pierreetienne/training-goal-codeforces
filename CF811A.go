package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF811A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF811A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF811A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF811A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF811A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF811A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF811A
Date: 31/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/811/A
Problem: CF811A
**/
func (in *CF811A) Run() {
	a, b := in.NextInt(), in.NextInt()
	c := 1
	for i := 0; a >= 0 && b >= 0; i++ {
		if i%2 == 0 {
			a -= c
			if a < 0 {
				fmt.Println("Vladik")
				break
			}

		} else {
			b -= c
			if b < 0 {
				fmt.Println("Valera")
			}

		}

		c++
	}

}

func NewCF811A(r *bufio.Reader) *CF811A {
	return &CF811A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF811A(bufio.NewReader(os.Stdin)).Run()
}
