package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprint(wr, x...)
}

func nextInt() int {
	i, e := strconv.Atoi(next())
	if e != nil {
		panic(e)
	}
	return i
}

func nextInt64() int64 {
	i, e := strconv.ParseInt(next(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	defer wr.Flush()
	sc.Buffer(make([]byte, 0), 10000009)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = nextInt()
		}
	}
	solve(n, m)

}

func solve(n int, m [][]int) {
	// Solution
	out("solution\n", m)
}
