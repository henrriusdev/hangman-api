package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Questions struct {
	ID           int
	Question     string
	Answer       string
	AnswerLenght int
}

func makeCN() (*sql.DB, error) {
	connection := "user=Henrry dbname=hangman password=postgres host=host.docker.internal port=4321 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	} else if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected")
	return db, nil
}

func GetQuestion(id int) []Questions {
	db, err := makeCN()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	query := "SELECT hint, answer, answer_lenght FROM quests WHERE id = $1"
	row := db.QueryRow(query, id)
	defer db.Close()
	if row.Err() != nil {
		log.Println(row.Err().Error())
		return nil
	}

	var quest Questions
	err = row.Scan(&quest.Question, &quest.Answer, &quest.AnswerLenght)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var hangman []Questions
	hangman = append(hangman, quest)

	return hangman
}

func Insert(hint string, answer string) error {
	db, err := makeCN()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	query := "INSERT INTO quests (hint, answer, answer_lenght) VALUES ($1,$2,$3)"

	_, err = db.Exec(query, hint, answer, len(answer))
	defer db.Close()
	if err != nil {
		return err
	}

	return nil
}
