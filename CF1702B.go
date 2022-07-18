package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1702B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1702B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1702B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1702B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1702B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1702B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1702B
Date: 16/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1702/B
Problem: CF1702B
**/
func (in *CF1702B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()

		ans := 1
		last := 0
		for last < len(str) {
			mapa := make(map[uint8]bool)
			for i := last; i < len(str); i++ {
				mapa[str[i]] = true
				if len(mapa) == 4 {
					ans++
					last = i
					break
				}
				if i + 1 == len(str){
					last = len(str)
				}
			}
		}

		fmt.Println(ans)
	}
}

func NewCF1702B(r *bufio.Reader) *CF1702B {
	return &CF1702B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1702B(bufio.NewReader(os.Stdin)).Run()
}
