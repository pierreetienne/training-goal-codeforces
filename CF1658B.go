package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1658B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1658B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1658B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1658B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1658B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1658B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Gopherbots hard number theory, math
Run solve the problem CF1658B
Date: 2/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1658/B
Problem: CF1658B
**/
func (in *CF1658B) Run() {
	const d = int64(998244353)
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		if n%2 == 0 {
			n /= 2
			ans := int64(1)
			for n > 1 {
				ans = ((n * n) % d * ans % d) % d
				n--
			}
			fmt.Println(ans)
		} else {
			fmt.Println(0)
		}
	}
}

func NewCF1658B(r *bufio.Reader) *CF1658B {
	return &CF1658B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1658B(bufio.NewReader(os.Stdin)).Run()
}
