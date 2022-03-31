package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF302A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF302A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF302A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF302A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF302A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF302A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF302A
Date: 29/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/302/A
Problem: CF302A
**/
func (in *CF302A) Run() {
	n, m := in.NextInt(), in.NextInt()

	ones := 0
	other := 0
	for i := 0; i < n; i++ {
		val := in.NextInt()
		if val == 1 {
			ones++
		} else {
			other++
		}
	}
	var ans strings.Builder
	for i := 0; i < m; i++ {
		l, r := in.NextInt(), in.NextInt()
		diff := (r - l) + 1
		if diff%2 == 0 && diff/2 <= ones && diff/2 <= other && diff > 1 {
			ans.WriteString("1\n")
		} else {
			ans.WriteString("0\n")
		}
	}
	fmt.Print(ans.String())

}

func NewCF302A(r *bufio.Reader) *CF302A {
	return &CF302A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF302A(bufio.NewReader(os.Stdin)).Run()
}
