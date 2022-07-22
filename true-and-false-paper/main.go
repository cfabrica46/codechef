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

	fmt.Scan(&attemps)

	for i := 0; i < attemps; i++ {
		trueAndFalsePaper()
	}
}

func trueAndFalsePaper() {
	in := bufio.NewReader(os.Stdin)

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

	k, err := strconv.Atoi(values[1])
	if err != nil {
		return
	}

	if n < k {
		return
	}

	fmt.Println(n - k)
}
