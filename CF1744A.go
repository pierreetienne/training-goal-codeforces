package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1744A struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1744A) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1744A) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1744A) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1744A) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1744A) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1744A
Date: 10/17/2022
User: wotan
URL: https://codeforces.com/problemset/problem/1744/A
Problem: CF1744A
**/
func (in *CF1744A) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n := in.NextInt()
    mapa := make(map[int][]int)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
      val := in.NextInt()
      if _, e := mapa[val]; !e {
        mapa[val] = []int{}
      }
      mapa[val] = append(mapa[val], i)
      arr[i] = val
    }

    str := in.NextString()
    ans := true
    for i := 0; i < n; i++ {
      p := mapa[arr[i]]
      x := uint8(0)
      for j := 0; j < len(p); j++ {
        if x == uint8(0) {
          x = str[p[j]]
        } else if x != str[p[j]] {
          ans = false
          break
        }
      }
    }

    if ans {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }
}

func NewCF1744A(r *bufio.Reader) *CF1744A {
  return &CF1744A{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1744A(bufio.NewReader(os.Stdin)).Run()
}
