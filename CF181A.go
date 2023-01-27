package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF181A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF181A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF181A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF181A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF181A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF181A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF181A
Date: 1/25/2023
User: wotan
URL: https://codeforces.com/problemset/problem/181/A
Problem: CF181A
**/
func (in *CF181A) Run() {
  n, m := in.NextInt(), in.NextInt()

  ansX, ansY := 0, 0
  for i := 0; i < n; i++ {
    arr := in.NextString()
    for j := 0; j < m; j++ {
      if arr[j] == '*' {
        ansX = ansX ^ i
        ansY = ansY ^ j // TODO PORQUE FUNCIONA ????

      }
    }
  }

  fmt.Println(ansX+1, ansY+1)
}

func NewCF181A(r *bufio.Reader) *CF181A {
  return &CF181A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF181A(bufio.NewReader(os.Stdin)).Run()
}
