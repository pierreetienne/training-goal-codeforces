package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1762A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1762A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1762A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1762A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1762A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1762A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1762A
Date: 12/22/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1762/A
Problem: CF1762A
**/
func (in *CF1762A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		sum := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			sum += arr[i]
		}

		if sum%2 == 0 {
			fmt.Println(0)
		} else {
			min := 10000000
			for i := 0; i < n; i++ {
				index := i
				auxSum := sum
				auxIndex := arr[index]
				count := 0
				for sum%2 != 0 && arr[index] != 0 {
					sum -= arr[index]
					arr[index] = int(math.Floor(float64(arr[index]) / 2.))
					sum += arr[index]
					count++
				}

				if sum%2 == 0 {
					min = int(math.Min(float64(min), float64(count)))
				}
				sum = auxSum
				arr[index] = auxIndex
			}

			fmt.Println(min)

		}

	}
}

func NewCF1762A(r *bufio.Reader) *CF1762A {
	return &CF1762A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1762A(bufio.NewReader(os.Stdin)).Run()
}
