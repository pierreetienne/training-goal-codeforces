package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1649A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1649A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1649A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1649A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1649A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1649A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1649A
Date: 11/10/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1649/A
Problem: CF1649A
**/
func (in *CF1649A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		ans := 0

		last := n - 1
		for ; last >= 0; last-- {
			if arr[last] == 1 {
				continue
			} else {
				break
			}
		}
		last++
		ini := 0

		for ; ini <= last; ini++ {
			if arr[ini] == 1 {
				continue
			} else {
				break
			}
		}
		ini--

		start := ini
		zero := false
		for i := ini; i <= last; i++ {
			if arr[i] == 0 {
				zero = true
			}
			if zero && arr[i] == 1 {
				ans += i - start
				start = i
			}
		}

		fmt.Println(ans)

	}

}

func NewCF1649A(r *bufio.Reader) *CF1649A {
	return &CF1649A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1649A(bufio.NewReader(os.Stdin)).Run()
}
