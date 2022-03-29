package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type CF49A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF49A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF49A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF49A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF49A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF49A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF49A
Date: 28/03/22
User: pepradere
URL: https://codeforces.com/problemset/problem/49/A
Problem: CF49A
**/
func (in *CF49A) Run() {
	vol := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"y": true,
	}

	str := strings.ToLower(in.GetLine())
	for i := len(str) - 1; i >= 0; i-- {
		if _, ok := vol[string(str[i])]; ok {
			fmt.Println("YES")
			break
		} else if unicode.IsLetter(rune(str[i])) {
			fmt.Println("NO")
			break
		}
	}
}

func NewCF49A(r *bufio.Reader) *CF49A {
	return &CF49A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF49A(bufio.NewReader(os.Stdin)).Run()
}
