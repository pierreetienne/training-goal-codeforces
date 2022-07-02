package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1700B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1700B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1700B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1700B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1700B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1700B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots time solve 40 ms
Run solve the problem CF1700B
Date: 29/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1700/B
Problem: CF1700B
**/
func (in *CF1700B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()

		f := int(str[0]-'0') + 2

		if f >= 10 {
			f = 1
		}

		var temp strings.Builder

		carry := 0
		for i := n - 1; i >= 0; i-- {
			num := int(str[i]-'0') + carry
			if num > f {
				val := 0
				if i == 0 {
					val = (f * 11) - num
				} else {
					val = (f + 10) - num
				}
				carry = 1
				temp.WriteRune(rune(val+'0'))
			} else if f > num {
				val := f - num
				temp.WriteRune(rune(val+'0'))
				carry = 0
			} else {
				temp.WriteRune('0')
				carry = 0
			}
		}

		var sol strings.Builder
		str = temp.String()
		for i:=len(str)-1;i>=0;i--{
			sol.WriteRune(rune(str[i]))
		}

		fmt.Println(sol.String())
	}
}

func NewCF1700B(r *bufio.Reader) *CF1700B {
	return &CF1700B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1700B(bufio.NewReader(os.Stdin)).Run()
}
