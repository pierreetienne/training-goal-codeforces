package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1759B struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1759B) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1759B) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1759B) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1759B) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1759B) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1759B
Date: 11/20/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1759/B
Problem: CF1759B
**/
func (in *CF1759B) Run() {
  for t := in.NextInt(); t > 0; t-- {
    m, s := in.NextInt(), in.NextInt()
    mapa := make(map[int]bool)
    max := 0
    for i := 0; i < m; i++ {
      val := in.NextInt()
      mapa[val] = true
      if max < val {
        max = val
      }
    }
    sum := 0
    for i := 1; i < max; i++ {
      if _, e := mapa[i]; !e {
        sum += i
      }
    }
    for sum < s {
      sum += max + 1
      max++
    }

    if sum == s {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }

  }
}

func NewCF1759B(r *bufio.Reader) *CF1759B {
  return &CF1759B{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1759B(bufio.NewReader(os.Stdin)).Run()
}
