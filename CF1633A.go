package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1633A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1633A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1633A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1633A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1633A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1633A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1633A
Date: 7/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1633/A
Problem: CF1633A
**/
func (in *CF1633A) Run() {
	for t := in.NextInt(); t > 0; t-- {

		n := in.NextInt()
		ans := -1
		diff := 100000
		for i := n; i > 9; i-- {
			if i%7 == 0 {
				if ans != -1 {
					if diff > dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n)) {
						diff = dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n))
						ans = i
					}
				} else {
					diff = dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n))
					ans = i
				}
			}
		}

		for i := n + 1; i < 5000; i++ {
			if i%7 == 0 {
				if ans != -1 {
					if diff > dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n)) {
						diff = dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n))
						ans = i
					}
				} else {
					diff = dif(fmt.Sprintf("%d", i), fmt.Sprintf("%d", n))
					ans = i
				}

			}
		}
		fmt.Println(ans)
	}
}

func dif(a, b string) int {
	ans := 0

	for i := 0; i < int(math.Min(float64(len(a)), float64(len(b)))); i++ {
		if a[i] != b[i] {
			ans++
		}
	}

	if len(a)-len(b) != 0 {
		ans += int(math.Abs(float64(len(a) - len(b))))
	}

	return ans
}

func NewCF1633A(r *bufio.Reader) *CF1633A {
	return &CF1633A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1633A(bufio.NewReader(os.Stdin)).Run()
}
