package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF931A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF931A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF931A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF931A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF931A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF931A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF931A
Date: 11/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/931/A
Problem: CF931A
**/
func (in *CF931A) Run() {
	a := in.NextInt()
	b := in.NextInt()

	diff := int(math.Abs(float64(a)-float64(b)))
	ans := 0

	if diff % 2 == 0 {
		div := diff/2
		sum := (div*(div+1))/2
		ans = 2*sum
	}else {
		div := int(math.Floor(float64(diff) / 2.))
		ans = (div*(div+1))/2
		ans += ((div+1)*(div+2))/2
	}

	fmt.Println(ans)
}

func NewCF931A(r *bufio.Reader) *CF931A {
	return &CF931A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF931A(bufio.NewReader(os.Stdin)).Run()
}
