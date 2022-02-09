package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1633B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1633B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1633B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1633B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1633B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1633B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1633B
Date: 8/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1633/B
Problem: CF1633B
**/
func (in *CF1633B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ones := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '1' {
				ones++
			}
		}
		zeros := len(str) - ones
		if ones == zeros {
			fmt.Println(ones - 1)
		} else {
			fmt.Println(int(math.Min(float64(zeros), float64(ones))))
		}
	}
}

func NewCF1633B(r *bufio.Reader) *CF1633B {
	return &CF1633B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1633B(bufio.NewReader(os.Stdin)).Run()
}
