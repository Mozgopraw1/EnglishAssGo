package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
)

type english struct{
	id int
	english_word string
	russian_word string
	time_init int
	time_check int
	examine int
	day int
}

var x int
var word string

func main() {
	openClose() // sqlite
	welcomeP() // приветствие и инструкция
	scanX(x) // выбор варианта продолжения программы
	variant(x) // запуск одного из вариантов
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
func scanX(x int) error{
	var a string
	_, err := fmt.Scan(a)
	x, err = strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err, x)
	if err != nil {
		return err
	}
	return err
}

// variant - запуск одного из вариантов
func variant(x int){
	if x == 1 {
		addWord()
	}
	if x == 2 {

	}
}

// addWord - добавление слова
func addWord(){
	fmt.Println("Введите слова по примеру:")
	fmt.Println("Word, Words")
	fmt.Println("Слово, Слова")
	scanWord(word)
	fmt.Printf(word)
	scanWord(word)
}

// scanWord - чтение слова из консоли
func scanWord(word string) {
	_, err := fmt.Fscanln(os.Stdin, word)
	if err != nil {
		panic(err)
	}
}