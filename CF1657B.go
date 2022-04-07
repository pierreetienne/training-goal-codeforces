package main

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type CF1657B struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1657B) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1657B) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *CF1657B) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1657B) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1657B) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1657B
Date: 5/04/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1657/B
Problem: CF1657B
**/
func (in *CF1657B) Run(){

}

func NewCF1657B(r *bufio.Reader) *CF1657B {
  return &CF1657B{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  NewCF1657B(bufio.NewReader(os.Stdin)).Run()
}