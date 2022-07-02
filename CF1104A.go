package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1104A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1104A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1104A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1104A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1104A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1104A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1104A
Date: 1/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1104/A
Problem: CF1104A
**/
func (in *CF1104A) Run() {
	n := in.NextInt()

	size := n
	num := 1

	var str strings.Builder
	str.WriteString(fmt.Sprintf("%d\n", size))
	for i := 0; i < size; i++ {
		if i == 0 {
			str.WriteString(fmt.Sprintf("%d", num))
		} else {
			str.WriteString(fmt.Sprintf(" %d", num))
		}

	}
	fmt.Println(str.String())

}

func NewCF1104A(r *bufio.Reader) *CF1104A {
	return &CF1104A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1104A(bufio.NewReader(os.Stdin)).Run()
}
