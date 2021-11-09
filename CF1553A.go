	package main

	import (
		"bufio"
		"fmt"
		"os"
		"strconv"
		"strings"
	)

	type CF1553A struct {
		sc        *bufio.Reader
		split     []string
		index     int
		separator string
	}

	func (in *CF1553A) GetLine() string {
		line, err := in.sc.ReadString('\n')
		if err != nil {
			fmt.Println("error line:", line, " err: ", err)
		}
		in.split = []string{}
		in.index = 0
		return line
	}
	func (in *CF1553A) load() {
		if len(in.split) <= in.index {
			in.split = strings.Split(in.GetLine(), in.separator)
			in.index = 0
		}
	}

	func (in *CF1553A) NextInt() int {
		in.load()
		val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
		in.index++
		return val
	}

	func (in *CF1553A) NextInt64() int64 {
		in.load()
		val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
		in.index++
		return val
	}

	func (in *CF1553A) NextString() string {
		in.load()
		val := strings.TrimSpace(in.split[in.index])
		in.index++
		return val
	}

	/**
	Run solve the problem CF1553A
	Date: 8/11/21
	User: pepradere
	URL: https://codeforces.com/problemset/problem/1553/A
	Problem: CF1553A
	**/
	func (in *CF1553A) Run() {
		t := in.NextInt()
		for ; t > 0; t-- {
			num := in.NextInt()
			str := fmt.Sprintf("%d", num)
			if str[len(str)-1] == '9' {
				num++
			}
			fmt.Println(num / 10)
		}
	}

	func NewCF1553A(r *bufio.Reader) *CF1553A {
		return &CF1553A{
			sc:        r,
			split:     []string{},
			index:     0,
			separator: " ",
		}
	}

	func main() {
		NewCF1553A(bufio.NewReader(os.Stdin)).Run()
	}
