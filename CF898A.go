package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF898A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF898A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF898A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF898A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF898A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF898A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF898A
Date: 25/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/898/A
Problem: CF898A
**/
func (in *CF898A) Run() {
	str := in.NextString()
	if str[len(str)-1] != 0 {
		ans,_ := strconv.Atoi(str)
		if str[len(str)-1] <= '5' {
			ans -= int(str[len(str)-1]-'0')
		}else {
			ans += 10-int(str[len(str)-1]-'0')
		}

		fmt.Println(ans)
	} else {
		fmt.Println(str)
	}

}

func NewCF898A(r *bufio.Reader) *CF898A {
	return &CF898A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF898A(bufio.NewReader(os.Stdin)).Run()
}
