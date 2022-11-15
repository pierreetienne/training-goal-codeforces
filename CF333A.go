package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF333A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF333A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF333A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF333A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF333A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF333A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF333A
Date: 11/4/2022
User: wotan
URL: https://codeforces.com/contest/333/problem/A
Problem: CF333A
**/
func (in *CF333A) Run() {
	n := in.NextInt64()

	p := int64(1)
	ans := int64(1)
	for p*int64(3) <= n {
		p *= int64(3)
	}
	r := p

	for r < n {
		r += p
		ans++
	}

	//	pp := float64(int64(math.Log(float64(n)) / math.Log(float64(3))))

	x := n
	for x%3 == 0 {
		x /= 3
	}
	//int64(math.Ceil(float64(n)/math.Pow(float64(3), pp))), " > ",
	fmt.Println((x / 3) + 1)

}

func NewCF333A(r *bufio.Reader) *CF333A {
	return &CF333A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF333A(bufio.NewReader(os.Stdin)).Run()
}
