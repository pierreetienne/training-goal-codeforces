package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1692C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1692C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1692C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1692C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1692C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1692C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Gopherbots
Run solve the problem CF1692C
Date: 16/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1692/C
Problem: CF1692C
**/
func (in *CF1692C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		in.NextString()
		arr := make([]string, 8)
		for i := 0; i < len(arr); i++ {
			arr[i] = in.NextString()
		}

		for i := 2; i < 8; i++ {
			for j := 2; j < 8; j++ {
				if arr[i-2][j-2] == '#' && arr[i-1][j-1] == '#' && arr[i][j] == '#' && arr[i-2][j] == '#' && arr[i][j-2] == '#' {
					fmt.Println(i, j)
					break
				}
			}
		}

	}
}

func NewCF1692C(r *bufio.Reader) *CF1692C {
	return &CF1692C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1692C(bufio.NewReader(os.Stdin)).Run()
}
