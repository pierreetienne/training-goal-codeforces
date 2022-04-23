package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1602A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1602A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1602A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1602A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1602A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1602A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Opcional para gopherbots
Run solve the problem CF1602A
Date: 21/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1602/A
Problem: CF1602A
**/
func (in *CF1602A) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		min := str[0]
		for i := 0; i < len(str); i++ {
			if str[i] < min {
				min = str[i]
			}
		}
		ans.WriteString(fmt.Sprintf("%s ", string(min)))
		found := false
		for i := 0; i < len(str); i++ {
			if str[i] == min && !found {
				found = true
				continue
			} else {
				ans.WriteString(string(str[i]))
			}
		}
		ans.WriteString("\n")
	}
	fmt.Print(ans.String())
}

func NewCF1602A(r *bufio.Reader) *CF1602A {
	return &CF1602A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1602A(bufio.NewReader(os.Stdin)).Run()
}
