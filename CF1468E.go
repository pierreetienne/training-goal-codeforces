package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1468E struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1468E) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1468E) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1468E) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1468E) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1468E) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1468E
Date: 1/5/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1468/E
Problem: CF1468E
**/
func (in *CF1468E) Run() {
	arr := make([]int, 4)
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		arr[0], arr[1], arr[2], arr[3] = in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
		sort.Ints(arr)
		ans.WriteString(fmt.Sprintf("%v\n", arr[0]*arr[2]))
	}
	fmt.Print(ans.String())
}

func NewCF1468E(r *bufio.Reader) *CF1468E {
	return &CF1468E{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1468E(bufio.NewReader(os.Stdin)).Run()
}
