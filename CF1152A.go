package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1152A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1152A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1152A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1152A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1152A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1152A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots time 60 ms
Run solve the problem CF1152A
Date: 23/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1152/A
Problem: CF1152A
**/
func (in *CF1152A) Run() {
	n, m := in.NextInt(), in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	indexOdd := 0
	indexEven := 0
	for i := 0; i < m; i++ {
		val := in.NextInt()
		if val%2 == 0 {
			indexEven++
		} else {
			indexOdd++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			if indexOdd > 0 {
				indexOdd--
				ans++
			}
		} else {
			if indexEven > 0 {
				indexEven--
				ans++
			}
		}
	}
	fmt.Println(ans)

}

func NewCF1152A(r *bufio.Reader) *CF1152A {
	return &CF1152A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1152A(bufio.NewReader(os.Stdin)).Run()
}
