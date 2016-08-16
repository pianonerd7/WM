package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "ANCWordFreq"
	DB_USER = "root"
	DB_PASS = "password"
)

type Word struct {
	Id        int
	Word      string
	Lemma     string
	POS       string
	Frequency int
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/ANCWordFreq?charset=utf8")
	checkErr(err)
	return db
}

func QueryByWord(word string) []Word {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM ANC WHERE Word='%s'", word)
	rows, err := db.Query(query)
	checkErr(err)

	var wordSlice []Word
	for rows.Next() {
		word := new(Word)
		err = rows.Scan(&word.Word, &word.Lemma, &word.POS, &word.Frequency, &word.Id)
		checkErr(err)
		fmt.Println(word)
		wordSlice = append(wordSlice, *word)
	}
	return wordSlice
}

func QueryByPOS(word string) []Word {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM ANC WHERE POS='%s'", word)
	rows, err := db.Query(query)
	checkErr(err)

	var wordSlice []Word
	for rows.Next() {
		word := new(Word)
		err = rows.Scan(&word.Word, &word.Lemma, &word.POS, &word.Frequency, &word.Id)
		checkErr(err)
		fmt.Println(word)
		wordSlice = append(wordSlice, *word)
	}
	return wordSlice
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
