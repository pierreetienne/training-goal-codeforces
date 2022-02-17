package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF796A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF796A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF796A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF796A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF796A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF796A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF796A
Date: 15/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/796/A
Problem: CF796A
**/
func (in *CF796A) Run() {
	n, m, k := in.NextInt(), in.NextInt(), in.NextInt()
	min := n * 10
	for i := 0; i < n; i++ {
		num := in.NextInt()
		if num != 0 && num <= k && (i+1) != m {
			aux := int(math.Abs(float64(m - (i + 1))))
			if aux < min {
				min = aux
			}
		}
	}
	fmt.Println(min*10)

}

func NewCF796A(r *bufio.Reader) *CF796A {
	return &CF796A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF796A(bufio.NewReader(os.Stdin)).Run()
}
