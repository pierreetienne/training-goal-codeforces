package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem A
Date: 24/04/22
User: pepradere
URL: CodeJamRoundB
Problem: A
**/

var aarr []int

func (in *A) Run() {
	ca := 1
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		aarr = make([]int, n)
		for i := 0; i < n; i++ {
			aarr[i] = in.NextInt()
		}

		ans := F(0, n-1, 0)

		fmt.Println(fmt.Sprintf("Case #%d: %d", ca, ans))
		ca++
	}
}

func F(i, j, max int) int {

	if i < len(arr) && j >= 0 && i <= j {
		sum := 0
		if aarr[i] <= max {
			sum = 1
		}

		valA := F(i+1, j, int(math.Max(float64(aarr[i]), float64(max)))) + sum
		fmt.Println("valA ", valA, i, j, max)

		sum = 0
		if aarr[j] <= max {
			sum = 1
		}

		valB := F(i, j-1, int(math.Max(float64(aarr[j]), float64(max)))) + sum
		fmt.Println("valB ", valB, i, j, max)

		if valA > valB {
			return valA
		}
		return valB

	}
	return 0

}

func NewA(r *bufio.Reader) *A {
	return &A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewA(bufio.NewReader(os.Stdin)).Run()
}
