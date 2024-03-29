package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	durationGame = 20
)

func main() {
	var attempts int

	fmt.Print("Enter the number of attempts: ")
	fmt.Scan(&attempts)

	for i := 0; i < attempts; i++ {
		chesstime()
	}
}

func chesstime() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of hours: ")

	line, err := in.ReadString('\n')
	if err != nil {
		return
	}

	line = line[:len(line)-1]

	duration, err := strconv.Atoi(line)
	if err != nil {
		return
	}

	duration = duration * 60

	fmt.Println(duration / durationGame)
}
