package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF1779B struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1779B) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1779B) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1779B) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1779B) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1779B) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1779B
Date: 3/01/23
User: pepradere
URL: https://codeforces.com/contests/1779/B
Problem: CF1779B
**/
func (in *CF1779B) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n := in.NextInt()
		if n == 2 {
			fmt.Println("YES")
			fmt.Println("9 5")
		} else {

			arr := make([]int, n)

			if n%2 == 0 {
				for i := 0; i < n; i++ {
					if i%2 == 0 {
						arr[i] = 1
					} else {
						arr[i] = -1
					}
				}
			} else {
				a := int(math.Floor(float64(n) / 2))
				for i := 0; i < n; i++ {
					if i%2 == 0 {
						arr[i] = -(a - 1)
					} else {
						arr[i] = a
					}
				}
			}

			ans := true
			sum := 0
			var sol strings.Builder
			for i := 0; i < n; i++ {
				sum += arr[i]
				sol.WriteString(fmt.Sprintf("%d ", arr[i]))
			}

			for i := 0; i < n-1; i++ {
				if arr[i]+arr[i+1] != sum || arr[i] == 0 {
					//	fmt.Println("sum ", sum, arr[i], " suma vecinos  ", arr[i]+arr[i+1])
					ans = false
					break
				}
			}

			if ans {
				fmt.Println("YES")
				fmt.Println(sol.String())
			} else {
				fmt.Println("NO")
			}

		}
	}
}

func NewCF1779B(r *bufio.Reader) *CF1779B {
	return &CF1779B{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1779B(bufio.NewReader(os.Stdin)).Run()
}
