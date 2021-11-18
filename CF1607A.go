package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1607A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1607A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1607A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1607A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1607A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1607A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1607A
Date: 17/11/21
User: pepradere
URL: https://codeforces.com/problemset/problem/1607/A
Problem: CF1607A
**/
func (in *CF1607A) Run() {
	t := in.NextInt()
	for ; t > 0; t-- {
		str := in.NextString()
		mapa := make(map[uint8]int)
		for i := 0; i < len(str); i++ {
			mapa[str[i]] = i
		}
		line := in.NextString()
		ans := 0
		for i := 1; i < len(line); i++ {
			ans += int(math.Abs(float64(mapa[line[i]] - mapa[line[i-1]])))
		}
		fmt.Println(ans)
	}
}

func NewCF1607A(r *bufio.Reader) *CF1607A {
	return &CF1607A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1607A(bufio.NewReader(os.Stdin)).Run()
}
