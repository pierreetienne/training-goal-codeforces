package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1764B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1764B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1764B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1764B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1764B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1764B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
APRENDIDO
GCD es el numero mas pequeño que se puede expreasar como una combinacion lineal

Run solve the problem CF1764B
Date: 2/9/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1764/B
Problem: CF1764B
**/
func (in *CF1764B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)

		tmp := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			tmp = gcd(tmp, arr[i])
		}

		fmt.Println(arr[n-1] / tmp)

	}
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func NewCF1764B(r *bufio.Reader) *CF1764B {
	return &CF1764B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1764B(bufio.NewReader(os.Stdin)).Run()
}
