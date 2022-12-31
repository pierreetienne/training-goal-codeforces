package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1654B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1654B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1654B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1654B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1654B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1654B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

var mapa map[uint8]int

/**
Run solve the problem CF1654B
Date: 9/18/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1654/B
Problem: CF1654B
**/
func (in *CF1654B) Run() {
	var ans strings.Builder
	for t := in.NextInt(); t > 0; t-- {
		mapa = map[uint8]int{}
		str := in.NextString()
		for i := 0; i < len(str); i++ {
			mapa[str[i]]++
		}
		for maxp := maxPrefix(str); maxp != -1; {
			str = str[maxp:]
			maxp = maxPrefix(str)

		}
		ans.WriteString(str + "\n")
	}
	fmt.Print(ans.String())
}

func maxPrefix(str string) int {
	pre := false
	str += "-"
	preMap := map[uint8]int{}
	for i := 1; i < len(str); i++ {
		preMap[str[i-1]]++
		if exist(preMap) {
			mapa[str[i-1]]--
			pre = true
			continue
		} else {
			if pre {
				return i - 1
			}
			break
		}
	}
	return -1
}

func exist(preMap map[uint8]int) bool {
	for k, v := range preMap {
		value, e := mapa[k]
		if !e {
			return false
		}
		if value <= v {
			return false
		}
	}
	return true
}

func NewCF1654B(r *bufio.Reader) *CF1654B {
	return &CF1654B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1654B(bufio.NewReader(os.Stdin)).Run()
}
