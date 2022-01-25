package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1539B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1539B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1539B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1539B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1539B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1539B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
// ADD GOPHERBOTS TEST CONTEST
Run solve the problem CF1539B
Date: 23/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1539/B
Problem: CF1539B
**/
func (in *CF1539B) Run() {
	n, q := in.NextInt(), in.NextInt()
	arr := make([]int, n+1)
	str := in.NextString()
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + int(str[i-1]-'a')
	}
	var sol strings.Builder
	for ; q > 0; q-- {
		l, r := in.NextInt(), in.NextInt()
		sol.WriteString(fmt.Sprintf("%d\n", arr[r]-arr[l-1]+(r-l)+1))
	}
	fmt.Print(sol.String())
}

func NewCF1539B(r *bufio.Reader) *CF1539B {
	return &CF1539B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1539B(bufio.NewReader(os.Stdin)).Run()
}
