package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CF14A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF14A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF14A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF14A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF14A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF14A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF14A
Date: 1/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/14/A
Problem: CF14A
**/
func (in *CF14A) Run() {
	n, _ := in.NextInt(), in.NextInt()
	str := make([]string, n)
	fI := 1000000
	fJ := 1000000
	lI := 0
	lJ := 0

	for i := 0; i < n; i++ {
		str[i] = in.NextString()
		for j := 0; j < len(str[i]); j++ {
			if str[i][j] == '*' {
				fI = int(math.Min(float64(fI), float64(i)))
				fJ = int(math.Min(float64(fJ), float64(j)))

				lI = int(math.Max(float64(lI), float64(i)))
				lJ = int(math.Max(float64(lJ), float64(j)))
			}
		}
	}

	var ans strings.Builder

	for i := fI; i <= lI; i++ {
		for j := fJ; j <= lJ; j++ {
			ans.WriteString(fmt.Sprintf("%v", string(str[i][j])))
		}
		ans.WriteString("\n")
	}

	fmt.Print(ans.String())

}

func NewCF14A(r *bufio.Reader) *CF14A {
	return &CF14A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF14A(bufio.NewReader(os.Stdin)).Run()
}
