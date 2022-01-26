package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1549B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1549B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1549B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1549B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1549B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1549B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1549B
Date: 24/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1549/B
Problem: CF1549B
**/
func (in *CF1549B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		strx := in.NextString()
		str2 := in.NextString()
		marcados := make([]int , n)
		for i:=0;i<n;i++ {
			if strx[i] == '1' {
				marcados[i] = 1
			}
		}
		ans := 0
		for i := 0; i < n; i++ {
			if str2[i] == '1' {
				if marcados[i] == 0 {
					ans++
					marcados[i] = 2
				} else {
					if i > 0  && marcados[i-1] == 1 {
						ans++
						marcados[i-1] = 2
					} else {
						if i+1 < n && marcados[i+1] == 1 {
							ans++
							marcados[i+1] = 2
						}
					}
				}
			}
		}
		fmt.Println(ans)
	}
}

func change(str string, pos int, char string) string {
	if pos == 0 {
		return char + str[1:]
	}
	if pos == len(str)-1 {
		return str[0:pos]+ char
	}
	return str[0:pos] + char + str[pos+1:]
}

func NewCF1549B(r *bufio.Reader) *CF1549B {
	return &CF1549B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1549B(bufio.NewReader(os.Stdin)).Run()
}
