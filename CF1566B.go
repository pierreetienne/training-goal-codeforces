package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1566B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1566B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1566B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1566B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1566B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1566B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1566B para gopherbots
Date: 29/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1566/B
Problem: CF1566B
**/
func (in *CF1566B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		str := in.NextString()
		ones := 0
		zeros := 0
		blocksZeros := 0
		tmp := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				zeros++
				tmp++
			} else {
				if tmp > 0 {
					blocksZeros++
					tmp = 0
				}
				ones++
			}
		}
		if tmp > 0 {
			blocksZeros++
			tmp = 0
		}

		if zeros == 0 {
			fmt.Println(0)
		} else if ones == 0 {
			fmt.Println(1)
		} else {

			if blocksZeros > 2 {
				fmt.Println(2)
			} else {
				fmt.Println(blocksZeros)
			}

		}

	}
}

func NewCF1566B(r *bufio.Reader) *CF1566B {
	return &CF1566B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1566B(bufio.NewReader(os.Stdin)).Run()
}
