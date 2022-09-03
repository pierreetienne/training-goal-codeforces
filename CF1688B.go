package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CF1688B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1688B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1688B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1688B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1688B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1688B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1688B
Date: 9/1/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1688/B
Problem: CF1688B
**/
func (in *CF1688B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int, n)
		odd := 0
		even := 0
		evenIndex := -1
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt()
			if arr[i]%2 == 0 {
				if evenIndex == -1 {
					evenIndex = i
				}
				if arr[evenIndex] > arr[i] {
					evenIndex = i
				}
				even++
			} else {
				odd++
			}
		}

		ans := 0
		if odd > 0 {
			ans = even
		} else {
			count := -1
			for i := 0; i < n; i++ {
				if arr[i]%2 == 0 {
					aux := 0
					for arr[i]%2 == 0 {
						arr[i] /= 2
						aux++
					}
					if count == -1 {
						count = aux
					}
					if aux < count {
						count = aux
					}
				}

			}
			ans = count + n - 1
		}
		fmt.Println(ans)
	}
}

func NewCF1688B(r *bufio.Reader) *CF1688B {
	return &CF1688B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1688B(bufio.NewReader(os.Stdin)).Run()
}
