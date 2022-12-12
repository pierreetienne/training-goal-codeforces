package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1131A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1131A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1131A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1131A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1131A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1131A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1131A
Date: 6/12/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1131/A
Problem: CF1131A
**/
func (in *CF1131A) Run() {
	w1, h1, w2, h2 := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
	ans := 2 * (h1 + h2 + 2)
	ans += int(math.Max(float64(w1), float64(w2))) * 2
	fmt.Println(ans)
}

func NewCF1131A(r *bufio.Reader) *CF1131A {
	return &CF1131A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1131A(bufio.NewReader(os.Stdin)).Run()
}
