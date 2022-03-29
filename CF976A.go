package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF976A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF976A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF976A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF976A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF976A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF976A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF976A
Date: 27/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/976/A
Problem: CF976A
**/
func (in *CF976A) Run() {
	n := in.NextInt()
	str := in.NextString()
	ones := 0
	zeroes := 0
	for i := 0; i < n; i++ {
		if str[i] == '1' {
			ones++
		} else {
			zeroes++
		}
	}

	ans := ""
	if ones > 0 {
		ans = "1"
	}
	if zeroes > 0 {
		for i := 0; i < zeroes; i++ {
			ans += "0"
		}
	}

	fmt.Println(ans)
}

func NewCF976A(r *bufio.Reader) *CF976A {
	return &CF976A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF976A(bufio.NewReader(os.Stdin)).Run()
}
