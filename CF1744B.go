package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1744B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1744B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1744B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1744B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1744B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1744B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**

Run solve the problem CF1744B
Date: 10/22/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1744/B
Problem: CF1744B
**/
func (in *CF1744B) Run() {
	var buff strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n, q := in.NextInt(), in.NextInt()
		sum := int64(0)
		odd, even := 0, 0
		for i := 0; i < n; i++ {
			val := in.NextInt64()
			if val%2 == 0 {
				even++
			} else {
				odd++
			}
			sum += val
		}

		for i := 0; i < q; i++ {
			ty, val := in.NextInt(), in.NextInt64()
			if ty == 0 {
				sum += val * int64(even)
				if val%2 != 0 {
					odd += even
					even = 0
				}
			} else {
				sum += val * int64(odd)
				if val%2 != 0 {
					even += odd
					odd = 0
				}
			}
			buff.WriteString(fmt.Sprintf("%v\n", sum))
		}
	}
	fmt.Print(buff.String())
}

func NewCF1744B(r *bufio.Reader) *CF1744B {
	return &CF1744B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1744B(bufio.NewReader(os.Stdin)).Run()
}
