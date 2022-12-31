package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	t := 0
	fmt.Scanln(&t)
	var str strings.Builder
	for ; t > 0; t-- {
		a, b, c := 0, 0, 0
		fmt.Scanln(&a, &b, &c)
		if a+b == c || a+c == b || b+c == a {
			str.WriteString("YES\n")
		} else {
			str.WriteString("NO\n")
		}
	}
	fmt.Print(str.String())

	//solve()
}

func solve() {
	sc := bufio.NewReader(os.Stdin)
	line, _ := sc.ReadString('\n')

	t, _ := strconv.Atoi(strings.TrimSpace(line))
	var str strings.Builder
	for ; t > 0; t-- {
		line, _ = sc.ReadString('\n')
		aux := strings.Split(line, " ")
		a, _ := strconv.Atoi(strings.TrimSpace(aux[0]))
		b, _ := strconv.Atoi(strings.TrimSpace(aux[1]))
		c, _ := strconv.Atoi(strings.TrimSpace(aux[2]))
		if a+b == c || a+c == b || b+c == a {
			str.WriteString("YES\n")
		} else {
			str.WriteString("NO\n")
		}

	}
	fmt.Print(str.String())
}
