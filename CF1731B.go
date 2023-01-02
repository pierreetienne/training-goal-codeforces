package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1731B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1731B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1731B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1731B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1731B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1731B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1731B
Date: 12/27/2022
User: wotan
URL: https://codeforces.com/contests/1731B
Problem: CF1731B
**/
func (in *CF1731B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		mod := int64(1000000000000)
		n := in.NextInt64() % mod

		a := int64(((n % mod) + int64(1)) % mod)
		b := int64(((int64(4) * n) % mod) - int64(1))

		c := ((n % mod) * (a % mod)) % mod
		e := (c % mod * b % mod) % mod
		d := e / int64(6)

		fmt.Println(((d % mod) * int64(2022)) % 1000000007)
	}
}

func NewCF1731B(r *bufio.Reader) *CF1731B {
	return &CF1731B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1731B(bufio.NewReader(os.Stdin)).Run()
}
