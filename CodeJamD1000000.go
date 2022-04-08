package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CodeJamD1000000 struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CodeJamD1000000) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CodeJamD1000000) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CodeJamD1000000) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CodeJamD1000000) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CodeJamD1000000) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CodeJamD1000000
Date: 2/04/22
User: pepradere
URL: https://codingcompetitions.withgoogle.com/codejam/round/0000000000876ff1/0000000000a46471
Problem: CodeJamD1000000
**/
func (in *CodeJamD1000000) Run() {
	c := 1
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
		}
		sort.Ints(arr)
		max := 0
		cont := 1
		for i := 0; i < n; i++ {
			if cont <= arr[i] {
				max++
				cont++
			}
		}
		newLine := "\n"
		if t-1 == 0 {
			newLine = ""
		}
		ans.WriteString(fmt.Sprintf("Case #%d: %d%s", c, max,newLine))
		c++
	}
	fmt.Print(ans.String())
}

func NewCodeJamD1000000(r *bufio.Reader) *CodeJamD1000000 {
	return &CodeJamD1000000{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCodeJamD1000000(bufio.NewReader(os.Stdin)).Run()
}
