package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1772B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1772B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1772B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1772B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1772B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1772B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1772B
Date: 12/21/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1772/B
Problem: CF1772B
**/
func (in *CF1772B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		arr := []int{
			in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt(),
		}
		arr[2], arr[3] = arr[3], arr[2]
		ws := false

		for i := 0; i < 4; i++ {
			if arr[i] < arr[(i+1)%4] && arr[i] < arr[(i+3)%4] && arr[(i+1)%4] < arr[(i+2)%4] && arr[(i+3)%4] < arr[(i+2)%4] {
				ws = true
			}
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1772B(r *bufio.Reader) *CF1772B {
	return &CF1772B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1772B(bufio.NewReader(os.Stdin)).Run()
}
