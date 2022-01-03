package main

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type CF1146A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1146A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1146A) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *CF1146A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1146A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1146A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1146A
Date: 2/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1146/A
Problem: CF1146A
**/
func (in *CF1146A) Run(){

}

func NewCF1146A(r *bufio.Reader) *CF1146A {
  return &CF1146A{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  NewCF1146A(bufio.NewReader(os.Stdin)).Run()
}