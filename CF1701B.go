package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1701B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1701B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1701B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1701B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1701B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1701B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots maybe could be hard 62 ms aprox to 100ms
Run solve the problem CF1701B
Date: 18/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1701/B
Problem: CF1701B
**/
func (in *CF1701B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		var ans strings.Builder

		arr := make([]bool, n+1)
		size := 2

		for i := 1; i <= n; i++ {
			if !arr[i] {
				for j := i; j <= n ; j*=size {
					ans.WriteString(fmt.Sprintf("%d ", j))
					arr[j] = true
				}
			}
		}

		fmt.Println(size)
		fmt.Println(ans.String())

	}
}

func NewCF1701B(r *bufio.Reader) *CF1701B {
	return &CF1701B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1701B(bufio.NewReader(os.Stdin)).Run()
}
