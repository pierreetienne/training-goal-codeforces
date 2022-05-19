package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1546A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1546A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1546A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1546A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1546A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1546A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1546A
Date: 17/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1546/A
Problem: CF1546A
**/
func (in *CF1546A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		a := make([]int, n)
		s := []int{}
		r := []int{}
		for i := 0; i < n; i++ {
			a[i] = in.NextInt()
		}
		b := make([]int, n)
		ans := 0
		for i := 0; i < n; i++ {
			b[i] = in.NextInt()
			if a[i]-b[i] > 0 {
				s = append(s, i)
			} else if a[i]-b[i] < 0 {
				r = append(r, i)
			}
			ans += a[i] - b[i]
		}
		if ans != 0 {
			fmt.Println(-1)
		} else {
			var sol strings.Builder
			ans := 0
			for i := 0; i < len(s); i++ {
				if a[s[i]] > b[s[i]] {
					for j := 0; j < len(r); j++ {
						if a[r[j]] < b[r[j]] {
							a[s[i]]--
							a[r[j]]++
							sol.WriteString(fmt.Sprintf("%d %d\n", s[i]+1, r[j]+1))
							ans ++
							if a[s[i]]== b[s[i]]{
								break
							}

							j = -1


						}
					}
					i=-1
				}
			}

			fmt.Println(ans)
			fmt.Print(sol.String())
		}

	}
}

func NewCF1546A(r *bufio.Reader) *CF1546A {
	return &CF1546A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1546A(bufio.NewReader(os.Stdin)).Run()
}
