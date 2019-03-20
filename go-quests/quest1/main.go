package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"time"
	"strconv"
)

type quiz struct {
	Problem string
	Answer string
}


func loadFile(filename string) (data []quiz){
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, line := range lines {
		data = append(data,quiz{Problem: line[0], Answer: line[1]})
	}

	return data
}

func timeout(data []quiz, countDownSec int) int{
	reader := bufio.NewReader(os.Stdin)
	var correct int
	ch := make(chan time.Duration, 1)
	go func() {
		start := time.Now()
		for _, question := range data {
			fmt.Println(question.Problem)
			text, _ := reader.ReadString('\n')
			if strings.TrimRight(text, "\n") == question.Answer {
				correct++
			}
		}
		ch <- time.Since(start)
	}()

	select {
	case elapsed := <-ch:
		fmt.Printf("All problem solved in %s\n", elapsed)
	case <- time.After(time.Duration(countDownSec) * time.Second):
		fmt.Println("Time's up")
	}
	return correct
}

func main() {
	countDownSec := 30
	if len(os.Args) > 1 {
		if arg, err := strconv.Atoi(os.Args[1]); err == nil {
			countDownSec = arg
		} else {
			fmt.Println("Usage:\nwith default count down time of 30s -> ./main ")
			fmt.Println("with specified count down time -> ./main time_in_sec")
			os.Exit(0)
		}
	}

	filename := "problems.csv"
	data := loadFile(filename)
	
	total := len(data)

	fmt.Println("Press Enter to start quiz.")
	reader := bufio.NewReader(os.Stdin)

	var correct int

	for {
		if input, _ := reader.ReadString('\n'); input == "\n"{
			correct = timeout(data, countDownSec)
			break
		}
	}

	fmt.Printf("There are %d questions in total. You got %d of them right.\n", total, correct)
}