package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1626A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1626A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1626A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1626A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1626A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1626A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1626A
Date: 21/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1626/A
Problem: CF1626A
**/
func (in *CF1626A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		mapa := make(map[uint8]int)
		var sb strings.Builder
		for i := 0; i < len(str); i++ {
			val := mapa[str[i]]
			mapa[str[i]] = val + 1
		}

		for k, v := range mapa {
			for i := 0; i < v; i++ {
				sb.WriteString(string(k))
			}
		}
		fmt.Println(sb.String())
	}
}

func NewCF1626A(r *bufio.Reader) *CF1626A {
	return &CF1626A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1626A(bufio.NewReader(os.Stdin)).Run()
}
