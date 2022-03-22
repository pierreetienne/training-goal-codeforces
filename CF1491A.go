package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1491A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1491A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1491A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1491A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1491A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1491A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1491A
Date: 18/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1491/A
Problem: CF1491A
**/
func (in *CF1491A) Run() {
	n, q := in.NextInt(), in.NextInt()

	ones := 0
	zeros := 0
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
		if arr[i] == 1 {
			ones++
		} else {
			zeros++
		}
	}
	var ans strings.Builder
	for i := 0; i < q; i++ {
		t := in.NextInt()
		x := in.NextInt()
		if t == 1 {
			value := 1 - arr[x-1]
			if value == 0 {
				zeros++
			} else {
				ones++
			}
			if arr[x-1] == 0 {
				zeros--
			} else {
				ones--
			}
			arr[x-1] = value
		} else {
			if x > ones {
				ans.WriteString("0\n")
			}else {
				ans.WriteString("1\n")
			}
		}
	}

	fmt.Print(ans.String())

}

func NewCF1491A(r *bufio.Reader) *CF1491A {
	return &CF1491A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1491A(bufio.NewReader(os.Stdin)).Run()
}
