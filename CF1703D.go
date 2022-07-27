package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1703D struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1703D) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1703D) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1703D) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1703D) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1703D) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots solve this proble brute force
Run solve the problem CF1703D
Date: 25/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1703/D
Problem: CF1703D
**/
func (in *CF1703D) Run() {
	var sol strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[string][]int)

		for i := 0; i < n; i++ {
			str := in.NextString()
			if _, e := mapa[str]; !e {
				mapa[str] = make([]int, 0)
			}
			arr := mapa[str]
			arr = append(arr, i)
			mapa[str] = arr
		}

		ans := make([]int, n)
		for s, val := range mapa {
			ws := false
			for j := 1; j < len(s); j++ {
				if _, e := mapa[s[0:j]]; e {
					if _, e := mapa[s[j:]]; e {
						ws = true
						break
					}
				}
			}
			for p := 0; p < len(val); p++ {
				if ws {
					ans[val[p]] = 1
				} else {
					ans[val[p]] = 0
				}

			}
		}

		for i := 0; i < n; i++ {
			sol.WriteString(fmt.Sprintf("%d", ans[i]))
		}
		sol.WriteString("\n")
	}
	fmt.Print(sol.String())
}

func NewCF1703D(r *bufio.Reader) *CF1703D {
	return &CF1703D{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1703D(bufio.NewReader(os.Stdin)).Run()
}
