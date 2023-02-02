package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF99A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF99A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF99A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF99A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF99A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF99A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF99A
Date: 28/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/99/A
Problem: CF99A
**/
func (in *CF99A) Run() {
	str := strings.Split(in.NextString(), ".")

	if str[0][len(str[0])-1] != '9' {

		if str[1][0] >= '5' {
			fmt.Println(str[0][0:len(str[0])-1] + string(rune(str[0][len(str[0])-1]+1)))
		} else {
			fmt.Println(str[0])
		}

	} else {
		fmt.Println("GOTO Vasilisa.")
	}
}

func NewCF99A(r *bufio.Reader) *CF99A {
	return &CF99A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF99A(bufio.NewReader(os.Stdin)).Run()
}
