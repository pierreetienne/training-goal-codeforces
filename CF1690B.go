package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1690B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1690B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1690B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1690B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1690B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1690B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Gopherbots
Run solve the problem CF1690B
Date: 10/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1690/B
Problem: CF1690B
**/
func (in *CF1690B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt64()
		}
		diff := int64(-1)
		ans := true
		b := make([]int64, n)
		for i := 0; i < n; i++ {
			b[i] = in.NextInt64()
			if a[i] >= b[i] {
				diff = int64(math.Max(float64(a[i])-float64(b[i]), float64(diff)))
			}
		}

		for i := 0; i < n; i++ {
			if a[i] < b[i] {
				ans = false
			} else if a[i]-diff < 0 && b[i] == 0 {
				continue
			} else if a[i]-diff == b[i] {
				continue
			} else {
				ans = false
			}

		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")

		}

	}
}

func NewCF1690B(r *bufio.Reader) *CF1690B {
	return &CF1690B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1690B(bufio.NewReader(os.Stdin)).Run()
}
