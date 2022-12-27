package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF363A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF363A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF363A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF363A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF363A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF363A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF363A
Date: 12/26/2022
User: wotan
URL: https://codeforces.com/problemset/problem/363/A
Problem: CF363A
**/
func (in *CF363A) Run() {
	str := in.NextString()
	var ans strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		val := int(str[i] - '0')
		if val == 0 {
			ans.WriteString("O-|-OOOO\n")
		} else if val == 1 {
			ans.WriteString("O-|O-OOO\n")
		} else if val == 2 {
			ans.WriteString("O-|OO-OO\n")
		} else if val == 3 {
			ans.WriteString("O-|OOO-O\n")
		} else if val == 4 {
			ans.WriteString("O-|OOOO-\n")
		} else if val == 5 {
			ans.WriteString("-O|-OOOO\n")
		} else if val == 6 {
			ans.WriteString("-O|O-OOO\n")
		} else if val == 7 {
			ans.WriteString("-O|OO-OO\n")
		} else if val == 8 {
			ans.WriteString("-O|OOO-O\n")
		} else if val == 9 {
			ans.WriteString("-O|OOOO-\n")
		}
	}
	fmt.Print(ans.String())
}

func NewCF363A(r *bufio.Reader) *CF363A {
	return &CF363A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF363A(bufio.NewReader(os.Stdin)).Run()
}
