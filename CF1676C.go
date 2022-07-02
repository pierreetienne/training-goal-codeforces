package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1676C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1676C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1676C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1676C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1676C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1676C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1676C
Date: 13/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1676/C
Problem: CF1676C
**/
func (in *CF1676C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextString()
		}
		min := -1
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					diff := 0
					for p := 0; p < m; p++ {
						d := int(math.Abs(float64(arr[i][p]) - float64(arr[j][p])))
						diff += d
					}
					if min == -1 {
						min = diff
					} else if min > diff {
						min = diff
					}

				}
			}
		}
		fmt.Println(min)
	}
}

func NewCF1676C(r *bufio.Reader) *CF1676C {
	return &CF1676C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1676C(bufio.NewReader(os.Stdin)).Run()
}
