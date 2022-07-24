package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var attempts int

	fmt.Print("Enter the number of attempts: ")
	fmt.Scan(&attempts)

	for i := 0; i < attempts; i++ {
		cricketRanking()
	}
}

func cricketRanking() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the stats of player A: ")

	r1, w1, c1, err := getStats(in)
	if err != nil {
		return
	}

	fmt.Print("Enter the stats of player B: ")

	r2, w2, c2, err := getStats(in)
	if err != nil {
		return
	}

	total1 := r1 + w1 + c1
	total2 := r2 + w2 + c2

	switch {
	case total1 > total2:
		fmt.Println("A")
	case total1 > total2:
		fmt.Println("B")
	case total1 == total2:
		fmt.Println("Its a draw")
	}
}

func getStats(in *bufio.Reader) (r, w, c int, err error) {
	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	line = line[:len(line)-1]

	values := strings.Split(line, " ")

	r, err = strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, 0, err
	}

	w, err = strconv.Atoi(values[1])
	if err != nil {
		return 0, 0, 0, err
	}

	c, err = strconv.Atoi(values[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return r, w, c, nil
}
