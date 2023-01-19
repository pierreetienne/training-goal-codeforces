package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF873A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF873A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF873A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF873A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF873A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF873A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF873A
Date: 12/01/23
User: pepradere
URL: https://codeforces.com/problemset/problem/873/A
Problem: CF873A
**/
func (in *CF873A) Run() {
  n, k, x := in.NextInt(), in.NextInt(), in.NextInt()

  sum := 0
  sumB := 0
  for i := 0; i < n; i++ {
    val := in.NextInt()
    sum += val
    if i+k >= n {
      sumB += val
    }
  }

  fmt.Println(sum - sumB + (k * x))

}

func NewCF873A(r *bufio.Reader) *CF873A {
  return &CF873A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF873A(bufio.NewReader(os.Stdin)).Run()
}
