package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CodeFamPunchedCards struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CodeFamPunchedCards) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CodeFamPunchedCards) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CodeFamPunchedCards) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CodeFamPunchedCards) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CodeFamPunchedCards) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CodeFamPunchedCards
Date: 2/04/22
User: pepradere
URL: https://codingcompetitions.withgoogle.com/codejam/round/0000000000876ff1/0000000000a4621b
Problem: CodeFamPunchedCards
**/
func (in *CodeFamPunchedCards) Run() {
	var ans strings.Builder
	c := 1
	for t := in.NextInt(); t > 0; t-- {
		R, C := in.NextInt(), in.NextInt()
		ans.WriteString(fmt.Sprintf("Case #%d:\n", c))
		c++

		for i := 0; i < (R*2)+1; i++ {
			for j := 0; j < (C*2)+1; j++ {
				if i <= 1 && j <= 1 {
					ans.WriteString(".")
				} else if i%2 == 0 {
					if j%2 == 0 {
						ans.WriteString("+")
					} else {
						ans.WriteString("-")
					}
				} else {
					if j%2 == 0 {
						ans.WriteString("|")
					} else {
						ans.WriteString(".")
					}
				}
			}
			ans.WriteString("\n")
		}

	}
	fmt.Print(ans.String())
}

func NewCodeFamPunchedCards(r *bufio.Reader) *CodeFamPunchedCards {
	return &CodeFamPunchedCards{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCodeFamPunchedCards(bufio.NewReader(os.Stdin)).Run()
}
