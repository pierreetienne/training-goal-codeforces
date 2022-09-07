package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1530B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1530B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1530B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1530B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1530B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1530B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Gopherbots max time 50 ms
Run solve the problem CF1530B
Date: 9/5/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1530/B
Problem: CF1530B
**/
func (in *CF1530B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		h, w := in.NextInt(), in.NextInt()

		m := make([][]int, h)
		for i := 0; i < h; i++ {
			m[i] = make([]int, w)
		}
		var str strings.Builder
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if valid(&m, i, j, h, w) {
					m[i][j] = 1
					str.WriteString("1")
				} else {
					str.WriteString("0")
				}

			}
			str.WriteString("\n")
		}
		fmt.Print(str.String())
	}
}

func valid(m *[][]int, a, b, h, w int) bool {
	dx := []int{1, -1, 0, 0, 1, -1, 1, -1}
	dy := []int{0, 0, 1, -1, 1, -1, -1, 1}

	if a >= 1 && b >= 1 && a < h-1 && b < w-1 {
		return false
	}
	count := 0
	for i := 0; i < len(dx); i++ {
		x, y := a+dx[i], b+dy[i]
		if x >= 0 && x < h && y >= 0 && y < w && (*m)[x][y] == 1 {
			count++
		}
	}
	return count == 0
}

func NewCF1530B(r *bufio.Reader) *CF1530B {
	return &CF1530B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1530B(bufio.NewReader(os.Stdin)).Run()
}
