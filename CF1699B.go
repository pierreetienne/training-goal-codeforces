package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type CF1699B struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CF1699B) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
    fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CF1699B) load() {
  if len(in.split) <= in.index {
    in.split = strings.Split(in.GetLine(), in.separator)
    in.index = 0
  }
}

func (in *CF1699B) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CF1699B) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CF1699B) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CF1699B
Date: 2/13/2023
User: wotan
URL: https://codeforces.com/problemset/problem/1699/B
Problem: CF1699B
**/
func (in *CF1699B) Run() {
  for t := in.NextInt(); t > 0; t-- {
    n, m := in.NextInt(), in.NextInt()
    matrix := make([][]int, n)
    for i := 0; i < n; i++ {
      matrix[i] = make([]int, m)
    }

    var ans strings.Builder
    par := []int{0, 1}
    impar := []int{1, 0}
    for i := 0; i < n; i++ {
      r := false
      for j := 0; j < m; j++ {
        if i%2 == 0 {
          if j%2 == 0 {
            if !r {
              matrix[i][j] = par[0]
            } else {
              matrix[i][j] = par[1]
            }
          } else {
            if !r {
              matrix[i][j] = par[1]
            } else {
              matrix[i][j] = par[0]
            }
            r = !r
          }
        } else {
          if j%2 == 0 {
            if !r {
              matrix[i][j] = impar[0]
            } else {
              matrix[i][j] = impar[1]
            }
          } else {
            if !r {
              matrix[i][j] = impar[1]
            } else {
              matrix[i][j] = impar[0]
            }
            r = !r
          }
        }
      }
      if i%2 != 0 {
        par, impar = impar, par
      }
      for j := 0; j < m; j++ {
        ans.WriteString(fmt.Sprintf("%d ", matrix[i][j]))
      }
      ans.WriteString("\n")
    }

    fmt.Print(ans.String())

  }
}

func NewCF1699B(r *bufio.Reader) *CF1699B {
  return &CF1699B{
    sc:        r,
    split:     []string{},
    index:     0,
    separator: " ",
  }
}

func main() {
  NewCF1699B(bufio.NewReader(os.Stdin)).Run()
}
