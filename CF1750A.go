package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1750A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1750A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1750A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1750A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1750A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1750A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1750A
Date: 11/15/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1750/A
Problem: CF1750A
**/
func (in *CF1750A) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n := in.NextInt()
    min := -1
    f := -1
    for i := 0; i < n; i++ {
      val := in.NextInt()
      if f == -1 {
        f = val
      }
      if min == -1 {
        min = val
      } else if min > val {
        min = val
      }
    }

    if f == min {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }

  }
}

func NewCF1750A(r *bufio.Reader) *CF1750A {
  return &CF1750A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1750A(bufio.NewReader(os.Stdin)).Run()
}
