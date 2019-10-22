package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type dbEng struct{
	id int
	englishWord string
	russianWord string
	timeInit time.Time
	timeCheck time.Time
	examine int
	day int
	word string
}

var x int

func main() {
	str := new(dbEng)
	openClose() // sqlite
	welcomeP() // приветствие и инструкция
	scanX(&x) // выбор варианта продолжения программы
	variant(x, *str) // запуск одного из вариантов
}

//openClose - sqlite
func openClose() {
	db, err := sql.Open("sqlite3", "english.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

//welcomeP - приветствие
func welcomeP() {
	fmt.Println("Введите цифру из списка для продолжения работы:")
	fmt.Println("1 - проверка слов")
	fmt.Println("2 - добавление слова")
}

//scanX - выбранные вариант продолжения программы.
func scanX(x *int) {
	var a string
	_, err := fmt.Scan(&a)
	*x, err = strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

// variant - запуск одного из вариантов
func variant(x int, str dbEng){
	if x == 1 {
		addWord(str)
	}
	if x == 2 {

	}
}

// addWord - добавление слова
func addWord(str dbEng) {
	fmt.Println("Введите слова по примеру:")
	fmt.Println("Word, Words")
	fmt.Println("Слово, Слова")
	scanWord(&str)
	str.word = str.englishWord
	scanWord(&str)
	str.word = str.russianWord

	//запись времени в переменные
	str.timeInit = time.Now()
	str.timeCheck = time.Now()

	//оставщиеся перменные
	str.examine = 1

	// day = 1 - Новое слово
	// day = 2 - 15 min
	// day = 3 - 2 hours
	// day = 4 - 1 day
	// day = 5 - 3 day
	// day = 6 - 7 day
	// day = 7 - 21 day
	// day = 8 - 50 day
	// day = 9 - 150 day
	// day = 10 - 365 day
	str.day = 1
}

// scanWord - чтение слова из консоли
func scanWord(str *dbEng) {
	_, err := fmt.Scan(&str.word)
	if err != nil {
		fmt.Println(err)
	}
}

/*
func addWordDB(str dbEng) {
	result, err := db.Exec(`insert into english (id, english_word, russian_word, time_int, time_check, examine, day) values()`)
if err != nil(
	panic(err))
}
*/