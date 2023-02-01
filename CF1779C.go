package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1779C struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1779C) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1779C) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1779C) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1779C) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1779C) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1779C
Date: 3/01/23
User: pepradere
URL:
Problem: CF1779C
**/
func (in *CF1779C) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n, m := in.NextInt(), in.NextInt()
    arr := make([]int, n)
    sumM := 0
    sum := 0
    for i := 0; i < n; i++ {
      arr[i] = in.NextInt()
      if i < m {
        sumM += arr[i]
        if i > 0 {
          arr[i] += arr[i-1]
        }
      }
      sum += arr[i]

    }

    ans := true

    for i := 0; i < m; i++ {
      if arr[i] > sum {
        ans = false
      }
    }

    if ans || n == 1 {
      fmt.Println(0)
    }
  }
}

func NewCF1779C(r *bufio.Reader) *CF1779C {
  return &CF1779C{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1779C(bufio.NewReader(os.Stdin)).Run()
}
