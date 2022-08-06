package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1714C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1714C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1714C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1714C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1714C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1714C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1714C
Date: 4/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1714/C
Problem: CF1714C
**/
func (in *CF1714C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		s := in.NextInt()
		ans := ""
		num := 9
		for {
			if s-num >= 0 {
				ans = fmt.Sprintf("%d%s", num, ans)
				s -= num
			}
			if s <= 0 {
				break
			}
			num--
		}
		fmt.Println(ans)

	}

}

func NewCF1714C(r *bufio.Reader) *CF1714C {
	return &CF1714C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1714C(bufio.NewReader(os.Stdin)).Run()
}
