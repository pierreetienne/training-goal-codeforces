package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1674B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1674B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1674B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1674B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1674B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1674B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1674B
Date: 10/05/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1674/B
Problem: CF1674B
**/
func (in *CF1674B) Run() {
	mapa := make(map[string]int)
	index := 1
	for i:= 'a'; i <= 'z'; i++{
		for j:='a';j<='z';j++{
			if i != j {
				mapa[fmt.Sprintf("%s%s", string(i), string(j))] = index
				index++
			}
		}
	}
	var str strings.Builder
	for t:=in.NextInt(); t > 0; t-- {
		str.WriteString(fmt.Sprintf("%d", mapa[in.NextString()]))
		str.WriteString("\n")
	}
	fmt.Print(str.String())
}

func NewCF1674B(r *bufio.Reader) *CF1674B {
	return &CF1674B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1674B(bufio.NewReader(os.Stdin)).Run()
}
