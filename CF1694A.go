package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1694A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1694A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1694A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1694A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1694A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1694A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots limit time  31 ms aprox 90ms
Run solve the problem CF1694A
Date: 20/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1694/A
Problem: CF1694A
**/
func (in *CF1694A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()
		var str strings.Builder
		for a+b > 0 {
			if a > 0 {
				str.WriteString("0")
				a--
			}
			if b > 0 {
				str.WriteString("1")
				b--
			}
		}
		fmt.Println(str.String())
	}

}

func NewCF1694A(r *bufio.Reader) *CF1694A {
	return &CF1694A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1694A(bufio.NewReader(os.Stdin)).Run()
}
