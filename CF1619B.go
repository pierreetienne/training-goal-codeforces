package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1619B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1619B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1619B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1619B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1619B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1619B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1619B
Date: 3/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1619/B
Problem: CF1619B
**/
func (in *CF1619B) Run() {
	mapa := make(map[int64]int)
	arr := make([]int64, 0)
	to := int64(math.Sqrt(1000000000))
	for i := int64(1); i <= to; i++ {
		if _, ok := mapa[i*i]; !ok {
			arr = append(arr, int64(i*i))
			mapa[i*i] = 1
		}
		if _, ok := mapa[i*i*i]; !ok {
			arr = append(arr, int64(i*i*i))
			mapa[i*i*i] = 1
		}
	}
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		count := 0
		for i := 0; i < len(arr); i++ {
			if arr[i] <= n {
				count++
			}
		}
		fmt.Println(count)
	}
}

func NewCF1619B(r *bufio.Reader) *CF1619B {
	return &CF1619B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1619B(bufio.NewReader(os.Stdin)).Run()
}
