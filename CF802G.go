package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF802G struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF802G) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF802G) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF802G) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF802G) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF802G) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF802G
Date: 21/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/802/G
Problem: CF802G
**/
func (in *CF802G) Run() {
	str := in.NextString()
	name := "heidi"
	pos := 0
	for i := 0; i < len(str); i++ {
		if str[i] == name[pos]  {
			pos++
			if pos == len(name) {
				break
			}
		}
	}
	if pos == len(name) && name != str {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func NewCF802G(r *bufio.Reader) *CF802G {
	return &CF802G{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF802G(bufio.NewReader(os.Stdin)).Run()
}
