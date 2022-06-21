package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1692B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1692B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1692B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1692B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1692B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1692B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots
Run solve the problem CF1692B
Date: 18/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1692/B
Problem: CF1692B
**/
func (in *CF1692B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[int]bool)
		count := 0
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if _, exist := mapa[val]; !exist {
				mapa[val] = true
			} else {
				count++
			}
		}
		if count%2 != 0 {
			fmt.Println(n - count - 1)
		} else {
			fmt.Println(n - count)
		}

	}
}

func NewCF1692B(r *bufio.Reader) *CF1692B {
	return &CF1692B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1692B(bufio.NewReader(os.Stdin)).Run()
}
