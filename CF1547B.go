package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1547B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1547B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1547B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1547B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1547B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1547B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1547B
Date: 18/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1547/B
Problem: CF1547B
**/
func (in *CF1547B) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		str := in.NextString()
		index := strings.Index(str, "a")
		ans := true
		if index == -1 {
			ans = false
		} else {
			letter := uint8('b')

			for left, right := index-1, index+1; ans ; letter++ {
				if left >= 0 {
					if str[left] == letter {
						left--
						continue
					}
				}
				if right < len(str) {
					if str[right] == letter {
						right++
						continue
					}
				}
				if left < 0 && right >= len(str){
					break
				}
				ans = false
			}
		}

		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func NewCF1547B(r *bufio.Reader) *CF1547B {
	return &CF1547B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1547B(bufio.NewReader(os.Stdin)).Run()
}
