package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1690C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1690C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1690C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1690C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1690C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1690C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1690C gopherbots
Date: 11/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1690/C
Problem: CF1690C
**/
func (in *CF1690C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}
		b := make([]int, n)
		for i := 0; i < n; i++ {
			b[i] = in.NextInt()
		}
		for i:=0;i<n;i++{
			if i == 0 {
				fmt.Print(b[i]-a[i])
			}else {
				fmt.Print(" ",b[i]-int(math.Max(float64(b[i-1]), float64(a[i]))))
			}
		}
		fmt.Println()
	}
}

func NewCF1690C(r *bufio.Reader) *CF1690C {
	return &CF1690C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1690C(bufio.NewReader(os.Stdin)).Run()
}
