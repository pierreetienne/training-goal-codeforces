package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1638A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1638A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1638A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1638A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1638A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1638A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1638A
Date: 12/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1638/A
Problem: CF1638A
**/
func (in *CF1638A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		index := 1
		ws := true
		pos := 0
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i] == i+1 && ws {
				index++
			} else {
				ws = false
			}
			if arr[i] == index {
				pos = i
			}
		}

		for i, j := index-1, pos; i < j; i++ {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}

		var sb strings.Builder
		for i := 0; i < n; i++ {
			sb.WriteString(fmt.Sprintf("%d ", arr[i]))
		}
		fmt.Println(sb.String())
	}
}

func NewCF1638A(r *bufio.Reader) *CF1638A {
	return &CF1638A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1638A(bufio.NewReader(os.Stdin)).Run()
}
