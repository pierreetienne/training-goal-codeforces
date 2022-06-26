package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF1617A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1617A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1617A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1617A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1617A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1617A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
gopherbots time 50ms
Run solve the problem CF1617A
Date: 24/06/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1617/A
Problem: CF1617A
**/
func (in *CF1617A) Run() {
	for t := in.NextInt(); t > 0; t-- {
		S := in.NextString()
		T := in.NextString()

		mapa := make(map[int]int)
		for i := 0; i < len(S); i++ {
			mapa[int(S[i])]++
		}

		key := make([]int, len(mapa))
		index := 0
		for k, _ := range mapa {
			key[index] = k
			index++
		}
		sort.Ints(key)

		ind := make([]int, 3)
		ind[0] = -1
		ind[1] = -1
		ind[2] = -1
		for i := 0; i < len(key); i++ {
			if key[i] == int(T[0]) {
				ind[0] = i
			}
			if key[i] == int(T[1]) {
				ind[1] = i
			}
			if key[i] == int(T[2]) {
				ind[2] = i
			}
		}

		if ind[1] < ind[2] && ind[1] != -1 && ind[2] != -1 && ind[0] != -1 && ind[0] < ind[1] {
			key[ind[1]], key[ind[2]] = key[ind[2]], key[ind[1]]
		}

		var str strings.Builder
		for i := 0; i < len(key); i++ {
			for j := 0; j < mapa[key[i]]; j++ {
				str.WriteString(string(rune(key[i])))
			}
		}

		fmt.Println(str.String())

	}
}

func NewCF1617A(r *bufio.Reader) *CF1617A {
	return &CF1617A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1617A(bufio.NewReader(os.Stdin)).Run()
}
