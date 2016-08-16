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
	"log"

	_ "github.com/ziutek/mymysql/godrv"
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
	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	return db
}

func UserById(id int) *Word {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT id, word, lemma, POS, Frequency FROM ANC WHERE id=?", id)
	word := new(Word)
	row.Scan(&word.Id, &word.Word, &word.Lemma, &word.POS, &word.Frequency)
	return word
}

func main() {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT Word, lemma, POS, Frequency, Id FROM ANC WHERE Id=1")
	word := new(Word)
	row.Scan(&word.Id, &word.Word, &word.Lemma, &word.POS, &word.Frequency)
	fmt.Println(word)
	//fmt.Println("id    : " + strconv.Itoa(user.Id) + "word  : " + word.Word + "frequency : \n" + word.Frequency)

}
