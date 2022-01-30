package main

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type CF1515A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1515A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1515A) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *CF1515A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1515A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1515A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1515A
Date: 28/01/22
User: pepradere
URL: https://codeforces.com/problemset/problem/1515/A
Problem: CF1515A
**/
func (in *CF1515A) Run(){

}

func NewCF1515A(r *bufio.Reader) *CF1515A {
  return &CF1515A{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  NewCF1515A(bufio.NewReader(os.Stdin)).Run()
}