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
	return wordSlice
}

func GetHighestFreqForWord(word string) Word {
	words := QueryByWord(word)

	if words.Len() == 0 {
		return *new(Word)
	}

	return *(words[0])
}

func QueryByPOS(pos string) Words {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM ANC WHERE POS='%s'", pos)
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
	return wordSlice
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("\nERROR\n")
		log.Fatal(err)
		panic(err)
	}
}

// GetUserWatermarkFromEmail takes an email and returns the watermark for that email
// if it exists
func GetUserWatermarkFromEmail(email string) string {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT secret FROM USERS WHERE email='%s'", email)
	rows, err := db.Query(query)
	checkErr(err)

	var watermark string
	for rows.Next() {
		err = rows.Scan(&watermark)
		checkErr(err)
	}
	return watermark
}

// GetUserEmailFromWaterMark takes an email and returns the watermark for that email
// if it exists
func GetUserEmailFromWaterMark(watermark string) string {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT email FROM USERS WHERE secret='%s'", watermark)
	rows, err := db.Query(query)
	checkErr(err)

	var email string
	for rows.Next() {
		err = rows.Scan(&email)
		checkErr(err)
	}
	return email
}

// QueryByUserWatermark takes a watermark and returns the email associted
// with that watermark if it exists
func CountUserWatermark(watermark string) int {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT COUNT(*) FROM USERS WHERE secret='%s'", watermark)
	rows, err := db.Query(query)
	checkErr(err)

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func CountUserEmail(email string) int {
	db := OpenDB()
	defer db.Close()

	query := fmt.Sprintf("SELECT COUNT(*) FROM USERS WHERE email='%s'", email)
	rows, err := db.Query(query)
	checkErr(err)

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}
	return count
}

// InsertNewUser takes an email and a watermark and adds a row to the
func InsertNewUser(email string, secret string) {
	db := OpenDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT USERS SET email=?, secret=?")
	checkErr(err)

	_, err = stmt.Exec(email, secret)
	checkErr(err)
}
