package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1255A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1255A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1255A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1255A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1255A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1255A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1255A
Date: 28/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1255/A
Problem: CF1255A
**/
func (in *CF1255A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		a, b := in.NextInt(), in.NextInt()

		diff := int(math.Abs(float64(a) - float64(b)))
		div := int(math.Floor(float64(diff) / 5.))
		ans := 0
		if div > 0 {
			diff -= div * 5
			ans += div
		}
		if diff > 0 {
			if diff == 4 {
				ans += 2
			}else if diff == 3 {
				ans+= 2
			}else if diff == 2 {
				ans++
			}else {
				ans++
			}
		}
		fmt.Println(ans)
	}
}

func NewCF1255A(r *bufio.Reader) *CF1255A {
	return &CF1255A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1255A(bufio.NewReader(os.Stdin)).Run()
}
