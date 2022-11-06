package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1670A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1670A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1670A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1670A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1670A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1670A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1670A
Date: 11/2/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1670/A
Problem: CF1670A
**/
func (in *CF1670A) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n := in.NextInt()
    arr := make([]int, n)

    for i := 0; i < n; i++ {
      val := in.NextInt()
      arr[i] = val
    }

    ans := true

    i, j := 0, 0
    for i < n {
      if arr[i] < 0 {
        arr[j] *= -1
        arr[i] *= -1
        j++
      }
      i++
    }

    sort := true
    for i := 1; i < n; i++ {
      if arr[i-1] > arr[i] {
        sort = false
      }
    }

    ans = sort

    if ans {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }

}

func NewCF1670A(r *bufio.Reader) *CF1670A {
  return &CF1670A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1670A(bufio.NewReader(os.Stdin)).Run()
}
