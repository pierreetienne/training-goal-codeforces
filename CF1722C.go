package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1722C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1722C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1722C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1722C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1722C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1722C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots ez time limit 100ms
Run solve the problem CF1722C
Date: 9/3/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1722/C
Problem: CF1722C
**/
func (in *CF1722C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapas := make([]map[string]bool, 3)

		for i := 0; i < 3; i++ {
			mapas[i] = make(map[string]bool)
		}

		sol := make([]int, 3)

		for i := 0; i < 3; i++ {
			for j := 0; j < n; j++ {
				str := in.NextString()
				mapas[i][str] = true
			}
		}

		for i := 0; i < 3; i++ {
			for str := range mapas[i] {
				points := 3
				count := 0
				for j := 0; j < 3; j++ {
					if j != i {
						_, e := mapas[j][str]
						if e {
							count++
						}
					}
				}
				if count == 2 {
					points = 0
				} else if count == 1 {
					points = 1
				}
				sol[i] += points
			}
		}

		fmt.Println(sol[0], sol[1], sol[2])

	}
}

func NewCF1722C(r *bufio.Reader) *CF1722C {
	return &CF1722C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1722C(bufio.NewReader(os.Stdin)).Run()
}
