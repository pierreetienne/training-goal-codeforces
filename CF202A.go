package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF202A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF202A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF202A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF202A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF202A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF202A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF202A
Date: 11/23/2022
User: wotan
URL: https://codeforces.com/problemset/problem/202/A
Problem: CF202A
**/
func (in *CF202A) Run() {
	str := in.NextString()
	arr := make([]int, len(str))
	for i := 0; i < len(arr); i++ {
		arr[i] = int(str[i])
	}
	sort.Ints(arr)
	var ans strings.Builder
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == arr[len(arr)-1] {
			ans.WriteString(fmt.Sprintf("%v", string(rune(arr[i]))))
		}
	}
	fmt.Println(ans.String())
}

func NewCF202A(r *bufio.Reader) *CF202A {
	return &CF202A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF202A(bufio.NewReader(os.Stdin)).Run()
}
