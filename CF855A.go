package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF855A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF855A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF855A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF855A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF855A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF855A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF855A
Date: 27/02/22
User: pepradere
URL: https://codeforces.com/problemset/problem/855/A
Problem: CF855A
**/
func (in *CF855A) Run() {
	n := in.NextInt()
	mapa := make(map[string]int)
	var ans strings.Builder
	for i := 0; i < n; i++ {
		val := in.NextString()
		if _, ok := mapa[val]; !ok {
			ans.WriteString("NO\n")
		} else {
			ans.WriteString("YES\n")
		}
		mapa[val] = 1
	}
	fmt.Print(ans.String())
}

func NewCF855A(r *bufio.Reader) *CF855A {
	return &CF855A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF855A(bufio.NewReader(os.Stdin)).Run()
}
