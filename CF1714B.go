package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1714B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1714B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1714B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1714B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1714B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1714B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots ez 100 ms
Run solve the problem CF1714B
Date: 3/08/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1714/B
Problem: CF1714B
**/
func (in *CF1714B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)

		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		mapa := make(map[int]bool)
		ans := 0
		for i := n - 1; i >= 0; i-- {
			_, e := mapa[arr[i]]
			if !e {
				mapa[arr[i]] = true
			} else {
				ans = i + 1
				break
			}
		}

		fmt.Println(ans)
	}
}

func NewCF1714B(r *bufio.Reader) *CF1714B {
	return &CF1714B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1714B(bufio.NewReader(os.Stdin)).Run()
}
