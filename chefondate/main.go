package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	nValues = 2
)

func main() {
	var attempts int

	fmt.Scan(&attempts)

	for i := 0; i < attempts; i++ {
		chefOnDate()
	}
}

func chefOnDate() {
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	line = line[:len(line)-1]

	values := strings.Split(line, " ")

	if len(values) != nValues {
		return
	}

	xValue, err := strconv.Atoi(values[0])
	if err != nil {
		return
	}

	yValue, err := strconv.Atoi(values[1])
	if err != nil {
		return
	}

	if xValue >= yValue {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
