package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1717A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1717A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1717A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1717A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1717A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1717A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1717A
Date: 15/09/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1717/A
Problem: CF1717A
**/
func (in *CF1717A) Run() {
	var ans strings.Builder
	arr := []int64{1, 4, 7, 10, 11, 16}
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		count := int64(0)

		div := int64(math.Floor(float64(n) / 6.))

		sum := int64(0)
		if n%6 != 0 {
			sum = arr[(n%6)-1]
		}
		count = (div * int64(16)) + sum

		fmt.Println(count)
	}
	fmt.Print(ans.String())
}

func gcd(a, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return (a / gcd(a, b)) * b
}

func NewCF1717A(r *bufio.Reader) *CF1717A {
	return &CF1717A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1717A(bufio.NewReader(os.Stdin)).Run()
}
