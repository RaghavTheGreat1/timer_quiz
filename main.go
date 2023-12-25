package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/RaghavTheGreat1/timer_quiz/models"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(file)
	}

	defer file.Close()

	problemsReasder := csv.NewReader(file)

	data, err := problemsReasder.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	parsedProblems := []models.Problem{}

	for _, v := range data {
		currentProblem := models.Problem{
			Question: v[0],
			Answer:   v[1],
		}

		parsedProblems = append(parsedProblems, currentProblem)

	}

	conductQuiz(parsedProblems)

}

func conductQuiz(problems []models.Problem) {
	timeLimit := flag.Int("Time Limit", 30, "Time Limit for the quiz")
	timer := time.NewTicker(time.Duration(*timeLimit) * time.Second)
	var marks int = 0

	for _, problem := range problems {
		fmt.Println("What is " + problem.Question + " ?")

		answerChannel := make(chan string)

		go func() {
			var guess string
			_, err := fmt.Scanf("%s", &guess)
			guess = strings.TrimSpace(guess)
			if err != nil {
				fmt.Println(err)
			}

			answerChannel <- guess

		}()

		select {
		case <-timer.C:

			fmt.Printf("Time's up, your total marks comes to: %d out of %d\n", marks, len(problems))
			return

		case answer := <-answerChannel:
			if answer == problem.Answer {
				marks++
			}
		}

	}

	fmt.Printf("That's it, your total marks comes to: %d out of %d\n", marks, len(problems))
}
