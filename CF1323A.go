package main

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type CF1323A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1323A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1323A) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *CF1323A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1323A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1323A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1323A
Date: 2/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1323/A
Problem: CF1323A
**/
func (in *CF1323A) Run(){

}

func NewCF1323A(r *bufio.Reader) *CF1323A {
  return &CF1323A{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  NewCF1323A(bufio.NewReader(os.Stdin)).Run()
}