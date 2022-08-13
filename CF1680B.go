package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1680B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1680B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1680B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1680B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1680B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1680B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots mm mid dificulty time 100 ms
Run solve the problem CF1680B
Date: 8/8/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1680/B
Problem: CF1680B
**/
func (in *CF1680B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, m := in.NextInt(), in.NextInt()

		pos := make([][]int, 0)

		for i := 0; i < n; i++ {
			data := in.GetLine()
			for j := 0; j < m; j++ {
				if data[j] == 'R' {
					pos = append(pos, []int{i, j, i + j})
				}
			}
		}

		sort.Slice(pos, func(i, j int) bool {
			if pos[i][2] < pos[i][2] {
				return true
			}
			return false
		})

		ws := true
		for i := 1; i < len(pos); i++ {
			if pos[i][0]-pos[0][0] < 0 || pos[i][1]-pos[0][1] < 0 {
				ws = false
			}
		}

		if ws {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1680B(r *bufio.Reader) *CF1680B {
	return &CF1680B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1680B(bufio.NewReader(os.Stdin)).Run()
}
