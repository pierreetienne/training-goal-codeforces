package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1159A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1159A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1159A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1159A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1159A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1159A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1159A
Date: 30/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1159/A
Problem: CF1159A
**/
func (in *CF1159A) Run() {
	n := in.NextInt()
	str := in.NextString()
	cont := 0
	for i:=0;i<n;i++{
		if str[i] == '-'{
			cont --
			if cont < 0 {
				cont = 0
			}
		}else {
			cont++
		}
	}
	fmt.Println(cont)
}

func NewCF1159A(r *bufio.Reader) *CF1159A {
	return &CF1159A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1159A(bufio.NewReader(os.Stdin)).Run()
}
