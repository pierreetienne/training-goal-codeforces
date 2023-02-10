package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1775A2 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1775A2) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1775A2) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1775A2) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1775A2) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1775A2) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1775A2
Date: 4/02/23
User: pepradere
URL: https://codeforces.com/problemset/problem/1775/A2
Problem: CF1775A2
**/
func (in *CF1775A2) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		a := string(str[0])
		b := str[1 : len(str)-1]
		c := string(str[len(str)-1])

		x, y := 1, 1
		//fmt.Println("a:", a, "b:", b, "c:", c, " ======")
		cont := 0
		for !((a >= b && b <= c) || (a <= b && b >= c)) {
			if cont == 0 {
				x++
				a = str[0:x]
				b = str[x : len(str)-y]
				c = str[len(str)-y:]
				//fmt.Println("a:", a, "b:", b, "c:", c, " << <")
				cont = 1
			} else if cont == 1 {
				x--
				y++
				a = str[0:x]
				b = str[x : len(str)-y]
				c = str[len(str)-y:]
				//fmt.Println("a:", a, "b:", b, "c:", c, " <++++++")
				cont = 2
			} else {
				cont = 0
			}
		}

		fmt.Println(a + " " + b + " " + c)
	}
}

func NewCF1775A2(r *bufio.Reader) *CF1775A2 {
	return &CF1775A2{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1775A2(bufio.NewReader(os.Stdin)).Run()
}
