package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1209A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1209A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1209A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1209A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1209A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1209A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1209A
Date: 12/4/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1209/A
Problem: CF1209A
**/
func (in *CF1209A) Run() {
	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}

	sort.Ints(arr)

	ans := 0
	for i := 0; i < n; i++ {
		if arr[i] != -1 {
			for j := 0; j < n; j++ {
				if arr[j] != -1 && arr[j]%arr[i] == 0 && i != j {

					arr[j] = -1
				}
			}
			arr[i] = -1
			ans++
		}
	}
	if n == 1 {
		ans = 1
	}
	fmt.Println(ans)
}

func NewCF1209A(r *bufio.Reader) *CF1209A {
	return &CF1209A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1209A(bufio.NewReader(os.Stdin)).Run()
}
