package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Questions struct {
	ID           int
	Question     string
	Answer       string
	AnswerLenght int
}

func makeCN() (*sql.DB, error) {
	connection := "user=postgres dbname=hangman password=Reyshell host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	} else if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected")
}

func GetQuestion(id int) []Questions {
	db, err := makeCN()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	query := "SELECT hint, answer, answer_lenght FROM quests WHERE id = $1"
	row := db.QueryRow(query, id)
	if row.Err() != nil {
		log.Fatal(err)
		return nil
	}

	var quest Questions
	err = row.Scan(&quest.ID, &quest.Question, &quest.Answer, &quest.AnswerLenght)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var hangman []Questions
	hangman = append(hangman, quest)

	return hangman
}

func Insert() error {
	return error
}
