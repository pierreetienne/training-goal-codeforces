package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF177A2 struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF177A2) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF177A2) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF177A2) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF177A2) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF177A2) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF177A2
Date: 1/21/2023
User: wotan
URL: https://codeforces.com/problemset/problem/177/A2
Problem: CF177A2
**/
func (in *CF177A2) Run() {
  n := in.NextInt()
  arr := make([][]int, n)
  sum := 0
  for i := 0; i < n; i++ {
    arr[i] = make([]int, n)
    for j := 0; j < n; j++ {
      arr[i][j] = in.NextInt()
      if i == (n-1)/2 {
        sum += arr[i][j]
        arr[i][j] = 0
      }
    }
  }

  for i := 0; i < n; i++ {
    sum += arr[i][i]
    arr[i][i] = 0
    sum += arr[i][n-i-1]
    arr[i][n-i-1] = 0
    if i == (n-1)/2 {
      for j := 0; j < n; j++ {
        sum += arr[j][i]
        arr[j][i] = 0
      }
    }
  }

  fmt.Println(sum)

}

func NewCF177A2(r *bufio.Reader) *CF177A2 {
  return &CF177A2{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF177A2(bufio.NewReader(os.Stdin)).Run()
}
