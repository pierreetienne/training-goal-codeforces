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

type CF1591C struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *CF1591C) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *CF1591C) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *CF1591C) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *CF1591C) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *CF1591C) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem CF1591C
Date: 12/12/21
User: pepradere
URL: https://codeforces.com/contest/1591/problem/C
Problem: CF1591C
**/
func (in *CF1591C) Run() {
	for t := in.NextInt(); t > 0; t-- {
		n, k := in.NextInt(), in.NextInt()
		arr := make([]int, n+1)
		center := 100
		arr[n] = center
		for i := 0; i < n; i++ {
			arr[i] = in.NextInt() + center
		}
		sort.Ints(arr)
		indexCenter := sort.SearchInts(arr, center)
		count := 1
		sum := 0
		current := center
		count, sum, current = F(&arr, 1, k, indexCenter, current, sum, count)
		count, sum, current = F(&arr, -1, k, indexCenter, current, sum, count)

		//min := int(math.Max(math.Abs(float64(arr[indexCenter]-arr[len(arr)-1])), math.Abs(float64(arr[indexCenter]-arr[0]))))

		//sum -= min
		fmt.Println(sum)

	}
}

func F(arr *[]int, dir int, k int, indexCenter int, current int, sum int, count int, start int) (int, int, int) {
	for i := indexCenter + dir; i >= 0 && i < len(*arr); i += dir {
		sum += int(math.Abs(float64(current - (*arr)[i])))
		fmt.Println(count%k, "sumo ", current, " menos ", (*arr)[i])
		if count%k == 0 {
			fmt.Println("entraaaaaaa")
			current = (*arr)[indexCenter]
			sum += int(math.Abs(float64((*arr)[indexCenter] - (*arr)[i])))
		} else {
			current = (*arr)[i]
		}
		count++
	}
	return count, sum, current
}

func NewCF1591C(r *bufio.Reader) *CF1591C {
	return &CF1591C{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewCF1591C(bufio.NewReader(os.Stdin)).Run()
}
