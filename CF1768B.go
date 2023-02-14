package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1768B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1768B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1768B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1768B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1768B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1768B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1768B
Date: 8/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1768/B
Problem: CF1768B
**/
func (in *CF1768B) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		find := 1
		for i := 0; i < n; i++ {
			val := in.NextInt()
			if val == find {
				find++
			}
		}

		diff := n - (find - 1)
		ans.WriteString(fmt.Sprintf("%d\n", int(math.Ceil(float64(diff)/float64(k)))))
	}
	fmt.Print(ans.String())
}

func NewCF1768B(r *bufio.Reader) *CF1768B {
	return &CF1768B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1768B(bufio.NewReader(os.Stdin)).Run()
}
