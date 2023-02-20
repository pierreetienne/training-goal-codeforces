package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1777B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1777B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1777B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1777B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1777B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1777B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1777B
Date: 18/02/23
User: pepradere
URL: https://codeforces.com/contest/1777/problem/B
Problem: CF1777B
**/
func (in *CF1777B) Run() {
	fact := make([]int64, 1e5+4)
	const MOD = int64(1e9 + 7)
	fact[0] = 1
	for i := int64(1); i < int64(len(fact)); i++ {
		fact[i] = ((fact[i-1] % MOD) * (i % MOD)) % MOD
	}

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		cont := ((((n - 1) % MOD) * (n % MOD)) % MOD * (fact[n] % MOD)) % MOD
		fmt.Println(cont)

	}
}

func NewCF1777B(r *bufio.Reader) *CF1777B {
	return &CF1777B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1777B(bufio.NewReader(os.Stdin)).Run()
}
