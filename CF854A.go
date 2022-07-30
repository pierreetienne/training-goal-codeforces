package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF854A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF854A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF854A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF854A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF854A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF854A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF854A
Date: 28/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/854/A
Problem: CF854A
**/
func (in *CF854A) Run() {
	n := in.NextInt()
	a, b := 0, 0
	if n%2 == 0 {
		a, b = (n/2)-1, (n/2)+1

	} else {
		a, b = (int)(math.Floor(float64(n)/float64(2))), n-(int)(math.Floor(float64(n)/float64(2)))
	}
	for i := 2; i < a+b; i++ {
		if a%i == 0 && b%i == 0 {
			a--
			b++
			break
		}
	}
	fmt.Println(a, b)
}

func NewCF854A(r *bufio.Reader) *CF854A {
	return &CF854A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF854A(bufio.NewReader(os.Stdin)).Run()
}
