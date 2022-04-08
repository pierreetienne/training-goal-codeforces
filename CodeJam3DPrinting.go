package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CodeJam3DPrinting struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CodeJam3DPrinting) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CodeJam3DPrinting) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CodeJam3DPrinting) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CodeJam3DPrinting) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CodeJam3DPrinting) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CodeJam3DPrinting
Date: 2/04/22
User: pepradere
URL: https://codingcompetitions.withgoogle.com/codejam/round/0000000000876ff1/0000000000a4672b
Problem: CodeJam3DPrinting
**/
func (in *CodeJam3DPrinting) Run() {
	var ans strings.Builder
	ca := 1
	for t := in.NextInt(); t > 0; t-- {
		ans.WriteString(fmt.Sprintf("Case #%d: ", ca))
		ca++
		a, b, c, d := 0, 0, 0, 0
		for i := 0; i < 3; i++ {
			xa, xb, xc, xd := in.NextInt(), in.NextInt(), in.NextInt(), in.NextInt()
			if i == 0 {
				a, b, c, d = xa, xb, xc, xd
			} else {
				if xa < a {
					a = xa
				}
				if xb < b {
					b = xb
				}
				if xc < c {
					c = xc
				}
				if xd < d {
					d = xd
				}
			}
		}

		if a > 0 && b > 0 && c > 0 && d > 0 && a+b+c+d >= 1000000 {
			arr := []int{a, b, c, d}
			index := 0
			for arr[0]+arr[1]+arr[2]+arr[3] > 1000000 {
				if arr[index] > 1 {
					arr[index]--
				} else {
					index++
				}
				if index > 3 {
					break
				}
			}

			ans.WriteString(fmt.Sprintf("%d %d %d %d\n", arr[0], arr[1], arr[2], arr[3]))
		} else {
			ans.WriteString("IMPOSSIBLE\n")
		}

	}
	fmt.Print(ans.String())
}

func NewCodeJam3DPrinting(r *bufio.Reader) *CodeJam3DPrinting {
	return &CodeJam3DPrinting{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCodeJam3DPrinting(bufio.NewReader(os.Stdin)).Run()
}
