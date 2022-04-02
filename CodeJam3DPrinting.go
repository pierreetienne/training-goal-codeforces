package main

import (
        "bufio"
        "fmt"
        "os"
        "sort"
        "strconv"
        "strings"
)
type CodeJam3DPrinting struct {
  sc        *bufio.Reader
  split     []string
  index     int
  separator string
}

func (in *CodeJam3DPrinting) GetLine() string {
  line, err := in.sc.ReadString('\n')
  if err != nil {
  fmt.Println("error line:", line, " err: ", err)
  }
  in.split = []string{}
  in.index = 0
  return line
}
func (in *CodeJam3DPrinting) load() {
  if len(in.split) <= in.index {
  in.split = strings.Split(in.GetLine(), in.separator)
  in.index = 0
  }
}

func (in *CodeJam3DPrinting) NextInt() int {
  in.load()
  val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
  in.index++
  return val
}

func (in *CodeJam3DPrinting) NextInt64() int64 {
  in.load()
  val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
  in.index++
  return val
}

func (in *CodeJam3DPrinting) NextString() string {
  in.load()
  val := strings.TrimSpace(in.split[in.index])
  in.index++
  return val
}

/**
Run solve the problem CodeJam3DPrinting
Date: 2/04/22
User: pepradere
URL: https://codingcompetitions.withgoogle.com/codejam/round/0000000000876ff1/0000000000a4672b
Problem: CodeJam3DPrinting
**/
func (in *CodeJam3DPrinting) Run(){

}

func NewCodeJam3DPrinting(r *bufio.Reader) *CodeJam3DPrinting {
  return &CodeJam3DPrinting{
  sc:        r,
  split:     []string{},
  index:     0,
  separator: " ",
  }
}

func main() {
  NewCodeJam3DPrinting(bufio.NewReader(os.Stdin)).Run()
}