package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF424A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF424A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF424A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF424A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF424A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF424A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF424A
Date: 5/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/424/A
Problem: CF424A
**/
func (in *CF424A) Run() {
	n := in.NextInt()
	str := in.NextString()
	u, d := 0, 0
	for i := 0; i < n; i++ {
		if str[i] == 'X' {
			u++
		} else {
			d++
		}
	}

	if u > d {
		diff := (n / 2) - d
		var ans strings.Builder

		c := 0
		for i := 0; i < n; i++ {
			if str[i] == 'X' && c < diff {
				ans.WriteString("x")
				c++
			} else {
				ans.WriteString(string(str[i]))
			}
		}
		fmt.Println(diff)
		fmt.Println(ans.String())
	} else if d > u {

		diff := (n / 2) - u
		var ans strings.Builder

		c := 0
		for i := 0; i < n; i++ {
			if str[i] == 'x' && c < diff {
				ans.WriteString("X")
				c++
			} else {
				ans.WriteString(string(str[i]))
			}
		}
		fmt.Println(diff)
		fmt.Println(ans.String())
	} else {
		fmt.Println(0)
		fmt.Println(str)
	}

}

func NewCF424A(r *bufio.Reader) *CF424A {
	return &CF424A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF424A(bufio.NewReader(os.Stdin)).Run()
}
