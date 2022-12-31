package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Testa struct {
	sc        *bufio.Reader
	split     []string
	index     int
	separator string
}

func (in *Testa) GetLine() string {
	line, err := in.sc.ReadString('\n')
	if err != nil {
		fmt.Println("error line:", line, " err: ", err)
	}
	in.split = []string{}
	in.index = 0
	return line
}
func (in *Testa) load() {
	if len(in.split) <= in.index {
		in.split = strings.Split(in.GetLine(), in.separator)
		in.index = 0
	}
}

func (in *Testa) NextInt() int {
	in.load()
	val, _ := strconv.Atoi(strings.TrimSpace(in.split[in.index]))
	in.index++
	return val
}

func (in *Testa) NextInt64() int64 {
	in.load()
	val, _ := strconv.ParseInt(strings.TrimSpace(in.split[in.index]), 10, 64)
	in.index++
	return val
}

func (in *Testa) NextString() string {
	in.load()
	val := strings.TrimSpace(in.split[in.index])
	in.index++
	return val
}

/**
Run solve the problem Testa
Date: 10/6/2022
User: wotan
URL: Test
Problem: Testa
**/
func (in *Testa) Run() {
	str := "codeforces"
	for i := 0; i < len(str); i++ {
		fmt.Println(int(str[i] - 'a'))
	}

}

func NewTesta(r *bufio.Reader) *Testa {
	return &Testa{
		sc:        r,
		split:     []string{},
		index:     0,
		separator: " ",
	}
}

func main() {
	NewTesta(bufio.NewReader(os.Stdin)).Run()
}
