package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1703B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1703B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1703B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1703B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1703B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1703B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**

Run solve the problem CF1703B
Date: 14/07/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1703/B
Problem: CF1703B
**/
func (in *CF1703B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		mapa := make(map[uint8]bool)

		str := in.NextString()
		ans := 0
		for i := 0; i < n; i++ {
			_, e := mapa[str[i]]

			if !e {
				mapa[str[i]] = true
				ans++
			}
			ans++
		}
		fmt.Println(ans)
	}

}

func NewCF1703B(r *bufio.Reader) *CF1703B {
	return &CF1703B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1703B(bufio.NewReader(os.Stdin)).Run()
}
