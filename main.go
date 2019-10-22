package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

type dbEng struct{
	id int
	englishWord string
	russianWord string
	timeInit int
	timeCheck int
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
}

// scanWord - чтение слова из консоли
func scanWord(str *dbEng) {
	_, err := fmt.Scan(&str.word)
	if err != nil {
		fmt.Println(err)
	}
}