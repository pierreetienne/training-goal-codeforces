package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CF769A struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF769A) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF769A) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF769A) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF769A) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF769A) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF769A
Date: 31/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/769/A
Problem: CF769A
**/
func (in *CF769A) Run() {

	puntos := [][]float64{
		{-10, 0},
		{-12, 8},
		{-4, 12},
		{0, 6},
		{2, 10},
		{16, 0},
	}

	//sqrt((x2-x1)^2 + (y2-y1)^2))
	sum := float64(0)
	for i := 1; i < len(puntos); i++ {
		a := math.Sqrt(((puntos[i][0] - puntos[i-1][0]) * (puntos[i][0] - puntos[i-1][0])) +
			((puntos[i][1] - puntos[i-1][1]) * (puntos[i][1] - puntos[i-1][1])))

		fmt.Println("distancia ", i, " es ", a, " x1 ", puntos[i-1][0], " x2 ", puntos[i][0])
		sum += a
	}

	fmt.Println(sum)

	n := in.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = in.NextInt()
	}
	sort.Ints(arr)
	if n == 1 {
		fmt.Println(arr[0])
	} else {
		fmt.Println(arr[(n-1)/2])
	}

	ma := make([][]float64, 10)
	for i := 0; i < 10; i++ {
		ma[i] = make([]float64, 10)
	}

}

func NewCF769A(r *bufio.Reader) *CF769A {
	return &CF769A{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF769A(bufio.NewReader(os.Stdin)).Run()
}
