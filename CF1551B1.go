package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1551B1 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1551B1) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1551B1) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1551B1) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1551B1) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1551B1) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1551B1
Date: 26/10/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1551/B1
Problem: CF1551B1
**/
func (in *CF1551B1) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		str := in.NextString()
		letters := make([]int, 32)
		colors := ""
		color := "R"
		r, g := 0, 0
		for i := 0; i < len(str); i++ {
			l := str[i]-'a'
			if letters[l] ==0 {
				colors+=color
				if color == "R" {
					r++
					color = "G"
				} else {
					g++
					color = "R"
				}
			} else if letters[l] == 1 {
				colors+=color
				if color == "R" {
					r++
					color = "G"
				} else {
					g++
					color = "R"
				}
			} else {
				colors+="W"
			}
			letters[l]++
		}

		ans := int(math.Min(float64(r), float64(g)))

		fmt.Println(ans)
	}

}

func NewCF1551B1(r *bufio.Reader) *CF1551B1 {
	return &CF1551B1{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1551B1(bufio.NewReader(os.Stdin)).Run()
}
