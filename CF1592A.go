package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1592A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1592A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1592A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1592A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1592A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1592A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Candidato gopherbots con tiempo de 110ms
Run solve the problem CF1592A
Date: 20/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1592/A
Problem: CF1592A
**/
func (in *CF1592A) Run() {
	var str strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n, H := in.NextInt(), in.NextInt()
		a, b := 0, 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if a < val {
				if b < a {
					b = a
				}
				a = val
			} else if b < val {
				b = val
			}
		}

		m := a + b
		x := int(math.Floor(float64(H / m)))

		diff := H - (x * m)

		ans := x * 2
		if diff > 0 {
			diff -= a
			ans++
		}
		if diff > 0 {
			diff -= b
			ans++
		}
		str.WriteString(fmt.Sprintf("%d\n", ans))
	}
	fmt.Print(str.String())
}

func NewCF1592A(r *bufio.Reader) *CF1592A {
	return &CF1592A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1592A(bufio.NewReader(os.Stdin)).Run()
}
