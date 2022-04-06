package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1536A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1536A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1536A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1536A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1536A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1536A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1536A
Date: 4/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1536/A
Problem: CF1536A
**/
func (in *CF1536A) Run() {

	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		arr := make([]int64, n)
		mapa := make(map[int64]bool)
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt64()
			mapa[arr[i]] = true
		}

		for i := 0; i < len(arr) && len(arr) <= 301; i++ {
			for j := 0; j < len(arr) && len(arr) <= 301; j++ {
				if i != j {
					mult := int64(math.Abs(float64(arr[i] - arr[j])))
					if _, ok := mapa[mult]; !ok {
						arr = append(arr, mult)
						mapa[mult] = true
					}
				}
			}
		}

		if len(arr) <= 300 {
			fmt.Println("YES")
			fmt.Println(len(arr))
			for i := 0; i < len(arr); i++ {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()
		} else {
			fmt.Println("NO")
		}

	}

}

func NewCF1536A(r *bufio.Reader) *CF1536A {
	return &CF1536A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1536A(bufio.NewReader(os.Stdin)).Run()
}
