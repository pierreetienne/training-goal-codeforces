package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF910A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF910A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF910A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF910A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF910A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF910A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF910A
Date: 12/12/2022
User: wotan
URL: https://codeforces.com/problemset/problem/910/A
Problem: CF910A
**/
func (in *CF910A) Run() {
	n, d := in.NextInt(), in.NextInt()
	str := in.NextString()

	ws := -1
	c := 0
	index := 0
	for {
		lastOne := -1
		f := false
		ini := index
		for j := 0; j <= d && index < n; index++ {
			if str[index] == '1' {
				lastOne = index
			}

			if index == n-1 && str[index] == '1' {
				f = true
				break
			}
			j++
		}

		if lastOne == -1 || ini == lastOne {
			ws = -1
			f = true
			break
		}
		index = lastOne
		c++
		ws = c
		if f {
			break
		}

	}

	fmt.Println(ws)

}

func NewCF910A(r *bufio.Reader) *CF910A {
	return &CF910A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF910A(bufio.NewReader(os.Stdin)).Run()
}
