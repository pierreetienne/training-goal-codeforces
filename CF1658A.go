package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1658A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1658A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1658A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1658A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1658A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1658A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1658A
Date: 19/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1658/A
Problem: CF1658A
**/
func (in *CF1658A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		in.NextInt()
		str := in.NextString()
		ans := 0

		for i := 1; i < len(str); i++ {
			if str[i] == '0' && str[i-1] == '0' {
				str = str[:i] + "11" + str[i:]
				ans += 2
			}
		}

		for i:=2;i< len(str);i++{
			if str[i-2]=='0' && str[i-1]=='1' && str[i] == '0'{
				str = str[:i-1] + "1" + str[i-1:]
				ans += 1
			}
		}

		fmt.Println(ans)
	}
}

func NewCF1658A(r *bufio.Reader) *CF1658A {
	return &CF1658A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1658A(bufio.NewReader(os.Stdin)).Run()
}
