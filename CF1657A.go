package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

type CF1657A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1657A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1657A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1657A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1657A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1657A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1657A
Date: 1/04/22
User: epradere
URL: https://codeforces.com/problemset/problem/1657/A
Problem: CF1657A
**/
func (in *CF1657A) Run() {
  for t := in.NextInt(); t > 0; t-- {
    x, y := in.NextInt(), in.NextInt()
    if x == 0 && y == 0 {
      fmt.Println(0)
    } else {
      dist := math.Sqrt(float64((x * x) + (y * y)))
      compare := float64(int(dist))
      if dist == compare {
        fmt.Println(1)
      } else {
        fmt.Println(2)
      }
    }
  }
}

func NewCF1657A(r *bufio.Reader) *CF1657A {
  return &CF1657A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1657A(bufio.NewReader(os.Stdin)).Run()
}