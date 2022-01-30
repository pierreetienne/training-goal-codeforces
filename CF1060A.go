package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1060A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1060A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1060A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1060A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1060A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1060A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1060A
Date: 29/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1060/A
Problem: CF1060A
**/
func (in *CF1060A) Run() {

	n := in.NextInt()
	str := in.NextString()
	ans := 0
	if n >= 11 {
		nums := make([]int, 10)
		for i := 0; i < n; i++ {
			nums[str[i]-'0']++
		}

		for {
			ws := false
			if nums[8] > 0 {
				count := 0
				nums[8]--
				for i := 0; i < 10; {
					if nums[i] > 0 && i != 8 {
						nums[i]--
						count++
					} else {
						i++
					}
					if count == 10 {
						ans++
						ws = true
						break
					}
				}
				if !ws && nums[8] >= (10-count) {
					nums[8] -= 10 - count
					ans++
					ws = true
				}
			}
			if !ws {
				break
			}
		}
		fmt.Println(ans)

	}else {
		fmt.Println(0)
	}
}

func NewCF1060A(r *bufio.Reader) *CF1060A {
	return &CF1060A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1060A(bufio.NewReader(os.Stdin)).Run()
}
