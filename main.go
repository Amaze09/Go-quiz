package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file")
	timeFlag := flag.Int("limit", 30, "Time limit for quiz")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Cannot open the file: %s", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Some problem parsing the csv file provided")
	}

	problems := parselines(lines)

	timer := time.NewTimer(time.Duration(*timeFlag) * time.Second)

	correct := 0
	
	for i, qe := range problems {
		fmt.Printf("Problem #%v %s =",i+1, qe.question)

		answerCh := make(chan string)

		go func ()  {
			var answer string
			fmt.Scanf("%s/n",&answer)
			answerCh <- answer
		}()
		
		select {
		case <- timer.C:
			fmt.Printf("\nYou got %d correct out of %d.\n", correct, len(problems))
			return
		
		case answer := <-answerCh:
			if answer == qe.answer {
				correct++
			}
		}

	}

}

	type problems struct {
		question string
		answer string
	}

	func parselines(lines [][]string) []problems {
		prob := make([]problems, len(lines))
		for i, line := range lines {
			prob[i] = problems{
				question : line[0],
				answer : strings.TrimSpace(line[1]),
			}
		}
		return prob
	}

	func exit(msg string)  {
		fmt.Println(msg)
		os.Exit(1)
	}