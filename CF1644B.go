package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1644B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1644B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1644B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1644B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1644B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1644B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1644B
Date: 16/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1644/B
Problem: CF1644B
**/
func (in *CF1644B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		var sb strings.Builder
		if n == 3 {
			sb.WriteString("3 2 1\n1 3 2\n3 1 2\n")
		} else {
			arr := make([]int, n)
			mapa := make(map[int]bool)
			for i := 0; i < n; i++ {
				mapa[i+1] = true
			}
			ws := true

			for ws {
				ws = false
				for i := 0; i < n; i++ {
					if i == 0 {
						arr[i] = 1
					} else if i == 1 {
						arr[i] = 2
					} else {
						arr[i] = Next(mapa, arr[i-1]+arr[i-2])
					}
					if arr[i] == -1 {
						ws = true
						for i := 0; i < n; i++ {
							mapa[i+1] = true
						}
						break
					}
					delete(mapa, arr[i])
				}

			}

			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					sb.WriteString(fmt.Sprintf("%d ", arr[(i+j)%n]))
				}
				sb.WriteString("\n")
			}
		}
		fmt.Print(sb.String())

	}
}

func Next(mapa map[int]bool, num int) int {
	for k, v := range mapa {
		if v && k != num {
			return k
		}
	}
	return -1
}

func NewCF1644B(r *bufio.Reader) *CF1644B {
	return &CF1644B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1644B(bufio.NewReader(os.Stdin)).Run()
}
