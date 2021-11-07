package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF731A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF731A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF731A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF731A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF731A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF731A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF731A
Date: 20/10/21
User: pepradere
URL: https://codeforces.com/contest/731/problem/A
Problem: CF731A
**/
func (in *CF731A) Run() {
	str := in.NextString()

	ans := 0
	var init uint8 = 'a'
	for i := 0; i < len(str); i++ {
		if init <= str[i] {
			ans += int(math.Min(float64(str[i]-init), float64(((init - 'a') + 1) + ('z' - str[i]))))
		} else {
			ans += int(math.Min(float64(init -str[i]), float64(((str[i] - 'a') + 1) + ('z' - init))))
		}
		init = str[i]
	}
	fmt.Println(ans)
}

func NewCF731A(r *bufio.Reader) *CF731A {
	return &CF731A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF731A(bufio.NewReader(os.Stdin)).Run()
}
