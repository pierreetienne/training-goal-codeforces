package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1523A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1523A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1523A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1523A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1523A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1523A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots simulation o(n*n) 171ms aprox to 250ms
Run solve the problem CF1523A
Date: 30/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1523/A
Problem: CF1523A
**/
func (in *CF1523A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		str := "0" + in.NextString() + "0"
		for m > 0 {
			var aux strings.Builder
			aux.WriteString("0")
			change := false
			for i := 1; i <= n; i++ {
				if str[i-1] == '0' && str[i] == '0' && str[i+1] == '1' {
					aux.WriteString("1")
					change = true
				} else if str[i-1] == '1' && str[i] == '0' && str[i+1] == '0' {
					aux.WriteString("1")
					change = true
				} else {
					aux.WriteString(string(str[i]))
				}
			}
			aux.WriteString("0")
			if change {
				str = aux.String()
			} else {
				break
			}

			m--
		}
		fmt.Println(str[1 : len(str)-1])
	}
}

func NewCF1523A(r *bufio.Reader) *CF1523A {
	return &CF1523A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1523A(bufio.NewReader(os.Stdin)).Run()
}
