package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1415A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1415A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1415A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1415A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1415A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1415A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Ez gopherbots solo hay que pensar como hallar los puntos mas lejanos del punto r,c
Run solve the problem CF1415A
Date: 7/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1415/A
Problem: CF1415A
**/
func (in *CF1415A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m, r, c := in.NextInt()-1, in.NextInt()-1, in.NextInt()-1, in.NextInt()-1

		pos := [][]int{{1, 0}, {0, 0}, {0, 1}, {1, 1}}
		max := 0
		for i := 0; i < len(pos); i++ {
			d := pos[i]
			x := d[0] * n
			y := d[1] * m

			val := int(math.Abs(float64(x)-float64(r))) + int(math.Abs(float64(y)-float64(c)))
			if max < val {
				max = val
			}

		}

		fmt.Println(max)

	}
}

func NewCF1415A(r *bufio.Reader) *CF1415A {
	return &CF1415A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1415A(bufio.NewReader(os.Stdin)).Run()
}
