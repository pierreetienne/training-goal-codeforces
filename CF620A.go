package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF620A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF620A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF620A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF620A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF620A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF620A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF620A
Date: 31/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/620/A
Problem: CF620A
**/
func (in *CF620A) Run() {

	x1, y1 := in.NextInt64(), in.NextInt64()
	x2, y2 := in.NextInt64(), in.NextInt64()
	max := int64(math.Max(math.Abs(float64(x1)-float64(x2)), math.Abs(float64(y1)-float64(y2))))
	fmt.Println(max)

}

func NewCF620A(r *bufio.Reader) *CF620A {
	return &CF620A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF620A(bufio.NewReader(os.Stdin)).Run()
}
