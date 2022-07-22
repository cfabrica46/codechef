package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var attemps int

	fmt.Print("Enter the number of attempts: ")
	fmt.Scan(&attemps)

	for i := 0; i < attemps; i++ {
		fillTheGrid()
	}
}

func fillTheGrid() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the value of M and N [M N]: ")
	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	line = line[:len(line)-1]

	values := strings.Split(line, " ")

	n, err := strconv.Atoi(values[0])
	if err != nil {
		return
	}

	m, err := strconv.Atoi(values[1])
	if err != nil {
		return
	}

	switch {
	case n%2 == 0 && m%2 == 0:
		fmt.Println(0)
	case n&2 != 0 && m%2 == 0:
		fmt.Println(m)
	case n&2 == 0 && m%2 != 0:
		fmt.Println(n)
	case n&2 != 0 && m%2 != 0:
		fmt.Println(m + n - 1)
	}
}
