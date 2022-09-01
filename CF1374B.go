package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1374B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1374B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1374B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1374B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1374B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1374B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1374B
Date: 8/26/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1374/B
Problem: CF1374B
**/
func (in *CF1374B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt64()
		mapa := map[int64]bool{
			n: true,
		}

		ans := 0
		count := 0
		for n != 1 {
			if n%6 == 0 {
				_, e := mapa[n/6]
				if e {
					ans = -1
					break
				}
				mapa[n/6] = true
				n /= 6
				count++
			} else {

				_, e := mapa[n*2]
				if e {
					ans = -1
					break
				}
				mapa[n*2] = true
				n *= 2
				count++
			}
		}

		if ans == -1 {
			fmt.Println(ans)
		} else {
			fmt.Println(count)
		}
	}
}

func NewCF1374B(r *bufio.Reader) *CF1374B {
	return &CF1374B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1374B(bufio.NewReader(os.Stdin)).Run()
}
