package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1691C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1691C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1691C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1691C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1691C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1691C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1691C
Date: 26/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1691/C
Problem: CF1691C
**/
func (in *CF1691C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt64()
		ones := make([]int, 0)

		str := in.NextString()
		for i := 0; i < n; i++ {
			val := str[i] - '0'
			if val == 1 {
				ones = append(ones, i)
			}
		}

		ans := "0"
		last := ones[len(ones)-1] + 1

		diff := n - last
		if int64(diff) > k {
			a := len(ones)
			b := a + (a / 10)
			ans = fmt.Sprintf("+%d%d", b, a)
		} else {
			a := len(ones)
			b := (a - 1) + (a / 10)
			ans = fmt.Sprintf("%d%d", b, a)
		}

		fmt.Println(ans)

	}
}

func NewCF1691C(r *bufio.Reader) *CF1691C {
	return &CF1691C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1691C(bufio.NewReader(os.Stdin)).Run()
}
