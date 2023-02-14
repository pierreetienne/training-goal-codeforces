package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1735B struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1735B) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1735B) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1735B) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1735B) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1735B) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1735B
Date: 2/11/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1735/B
Problem: CF1735B
**/
func (in *CF1735B) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n := in.NextInt()
    min := in.NextInt()
    cont := 0
    for i := 1; i < n; i++ {
      val := in.NextInt()
      if min*2 <= val {
        cont += (val - 1) / ((min * 2) - 1)
      }
    }

    fmt.Println(cont)
  }

}

func NewCF1735B(r *bufio.Reader) *CF1735B {
  return &CF1735B{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1735B(bufio.NewReader(os.Stdin)).Run()
}
