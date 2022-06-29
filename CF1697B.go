package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1697B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1697B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1697B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1697B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1697B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1697B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots resolvi en 234ms
Run solve the problem CF1697B
Date: 27/06/22
User: pepradere
URL:  https://codeforces.com/problemset/problem/1697/B
Problem: CF1697B
**/
func (in *CF1697B) Run() {
	n, q := in.NextInt(), in.NextInt()
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = int64(in.NextInt())
	}

	sort.Slice(sum, func(i, j int) bool {
		return sum[i] < sum[j]
	})

	for i := 1; i <= n; i++ {
		sum[i] += sum[i-1]
	}

	var ans strings.Builder
	for i := 0; i < q; i++ {
		x, y := in.NextInt(), in.NextInt()
		ans.WriteString(fmt.Sprintf("%d\n", sum[n-x+y]-sum[n-x]))
	}

	fmt.Print(ans.String())
}

func NewCF1697B(r *bufio.Reader) *CF1697B {
	return &CF1697B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1697B(bufio.NewReader(os.Stdin)).Run()
}
