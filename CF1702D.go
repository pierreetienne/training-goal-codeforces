package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1702D struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1702D) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1702D) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1702D) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1702D) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1702D) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
greedy gopherbots 100 ms
Run solve the problem CF1702D
Date: 27/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1702/D
Problem: CF1702D
**/
func (in *CF1702D) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		p := in.NextInt()
		m := make([][]int, 27)

		totalSum := 0
		for i := 0; i < len(str); i++ {
			val := int(str[i] - 'a')
			arr := m[val]
			if arr == nil {
				arr = make([]int, 0)
			}
			arr = append(arr, i+1)
			m[val] = arr
			totalSum += val + 1
		}

		if totalSum <= p {
			fmt.Println(str)
		} else {
			for i := len(m) - 1; i >= 0; i-- {
				if m[i] != nil {
					letter := i + 1
					ws := false
					for j := 0; j < len(m[i]); j++ {
						totalSum -= letter
						m[i][j] = -1

						if totalSum <= p {
							ws = true
							break
						}
					}
					if ws {
						break
					}
				}
			}

			ans := make([]uint8, len(str))
			for i := 0; i < len(m); i++ {
				if m[i] != nil {
					for j := 0; j < len(m[i]); j++ {
						if m[i][j] > 0 {
							ans[m[i][j]-1] = uint8(i + 'a')
						}
					}
				}
			}

			var solve strings.Builder
			for i := 0; i < len(ans); i++ {
				if ans[i] != 0 {
					solve.WriteString(string(ans[i]))
				}
			}
			fmt.Println(solve.String())
		}

	}
}

func NewCF1702D(r *bufio.Reader) *CF1702D {
	return &CF1702D{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1702D(bufio.NewReader(os.Stdin)).Run()
}
