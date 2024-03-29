package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1779A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1779A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1779A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1779A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1779A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1779A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1779A
Date: 3/01/23
User: pepradere
URL: https://codeforces.com/contests/1779/A
Problem: CF1779A
**/
func (in *CF1779A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		L := 0
		R := 0
		for i := 0; i < n; i++ {
			if str[i] == 'L' {
				L++
			} else {
				R++
			}
		}

		if L == n || R == n {
			fmt.Println(-1)
		} else if index := strings.Index(str, "LR"); index != -1 {
			fmt.Println(index + 1)
		} else if index := strings.Index(str, "RL"); index != -1 {
			fmt.Println(0)
		} else {
			fmt.Println(-1)
		}
	}
}

func NewCF1779A(r *bufio.Reader) *CF1779A {
	return &CF1779A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1779A(bufio.NewReader(os.Stdin)).Run()
}
