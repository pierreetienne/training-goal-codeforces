package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF411A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF411A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF411A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF411A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF411A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF411A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF411A
Date: 24/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/411/A
Problem: CF411A
**/
func (in *CF411A) Run() {
	str := in.NextString()

	a, b, c := false, false, false

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			a = true
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			b = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			c = true
		}
	}

	if a && b && c && len(str) >= 5 {
		fmt.Println("Correct")
	} else {
		fmt.Println("Too weak")
	}

}

func NewCF411A(r *bufio.Reader) *CF411A {
	return &CF411A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF411A(bufio.NewReader(os.Stdin)).Run()
}
