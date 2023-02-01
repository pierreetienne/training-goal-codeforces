package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ab struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *ab) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *ab) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *ab) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *ab) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *ab) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem ab
Date: 28/12/22
User: pepradere
URL: ab
Problem: ab
**/
func (in *ab) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		str := in.NextString()
		fmt.Println(solve(str, n))
	}
}

func solve(str string, n int) string {

	lastResult := int(str[0] - '0')
	//operationDiff := 0
	//operationAdd := 0
	var b strings.Builder

	for cont2 := 1; cont2 < n; cont2++ {
		element := int(str[cont2] - '0')
		//operationDiff = lastResult - element
		//operationAdd =

		if lastResult-element == -1 {
			lastResult = lastResult + element
			b.WriteString("+")
			continue
		}
		if lastResult-element < lastResult+element {
			lastResult = lastResult - element
			b.WriteString("-")
			continue
		}
		b.WriteString("+")
		lastResult = lastResult + element
	}
	return b.String()
}

func Newab(r *bufio.Reader) *ab {
	return &ab{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	Newab(bufio.NewReader(os.Stdin)).Run()
}
