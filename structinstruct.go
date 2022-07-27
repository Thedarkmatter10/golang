package main

import (
	"fmt"
)

type Alternative struct {
	Question
	AlternativeName string
}

type Question struct {
	Questionnaire
	QuestionName string
}

type Questionnaire struct {
	QuestionnaireName string
}

func main() {
	a := Alternative{
		Question: Question{
			Questionnaire: Questionnaire{
				QuestionnaireName: "q",
			},
		},
	}
	fmt.Printf("%v", a)
}
