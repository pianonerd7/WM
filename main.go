package main

//func main() {
//wordnet.InitWN()
//word := "It was a big, bold idea, especially given that Uber was still a relatively small start-up and no one in China had ever even heard of us.  And of course, anytime we got into a discussion about our efforts in China, most people thought we were naive, crazy – or both. We saw things differently of course. China is an inspiring country with astonishing opportunity. Many of the world’s mega cities are Chinese, and their thirst for transportation innovation is second to none. Uber’s mission to make “transportation as reliable as running water, everywhere for everyone” resonates especially strongly in China."
//fmt.Printf("Sense Results for '%s': %s\n", word, wordnet.FindTheInfo(word, 1, 5, 0))
//fmt.Printf("Sense Results for '%s': %s\n", word, wordnet.FindTheInfo_ds(word, 1, 5, 0))
//x := usubstitute.GetMapFromMessage(word)
//for key, value := range x {
//	fmt.Printf("%v , %v \n", key, value)
//}
//fmt.Println(watermark.EmbedMessage(word))
//usubstitute.GetMapFromMessage(word)
//}

import (
	"database/sql"
	"fmt"

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

///Users/annahe/src/code.uber.internal/go/src/github.com/Go-SQL-Driver/MySQL
func main() {
	// sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	//db, err := sql.Open("mysql", "reckhou:reckhou@/test?charset=utf8") user:password@/dbname
	db, err := sql.Open("mysql", "root:password@/ANCWordFreq?charset=utf8")
	checkErr(err)

	//Query
	rows, err := db.Query("SELECT * FROM ANC WHERE Id=2")
	checkErr(err)

	for rows.Next() {
		word := new(Word)
		err = rows.Scan(&word.Word, &word.Lemma, &word.POS, &word.Frequency, &word.Id)
		checkErr(err)
		fmt.Println(word)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
