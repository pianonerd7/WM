package sql

import (
	"database/sql"
	"fmt"
	"log"
	"sort"

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

type Words []*Word

func (slice Words) Len() int {
	return len(slice)
}

func (slice Words) Less(i, j int) bool {
	return slice[i].Frequency < slice[j].Frequency
}

func (slice Words) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type ByFrequency struct{ Words }

func (s ByFrequency) Less(i, j int) bool {
	return s.Words[i].Frequency >= s.Words[j].Frequency
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/ANCWordFreq?charset=utf8")
	checkErr(err)
	return db
}

func QueryByWord(word string) Words {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM ANC WHERE Word='%s'", word)
	rows, err := db.Query(query)
	checkErr(err)

	var wordSlice Words
	for rows.Next() {
		word := new(Word)
		err = rows.Scan(&word.Word, &word.Lemma, &word.POS, &word.Frequency, &word.Id)
		checkErr(err)
		wordSlice = append(wordSlice, word)
	}

	sort.Sort(ByFrequency{wordSlice})
	/*
		for _, o := range wordSlice {
			fmt.Println(o)
		}*/
	return wordSlice
}

func QueryByPOS(word string) Words {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM ANC WHERE POS='%s'", word)
	rows, err := db.Query(query)
	checkErr(err)

	var wordSlice Words
	for rows.Next() {
		word := new(Word)
		err = rows.Scan(&word.Word, &word.Lemma, &word.POS, &word.Frequency, &word.Id)
		checkErr(err)
		wordSlice = append(wordSlice, word)
	}

	sort.Sort(ByFrequency{wordSlice})
	/*
		for _, o := range wordSlice {
			fmt.Println(o)
		} */
	return wordSlice
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
