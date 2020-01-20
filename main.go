package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
	"time"
)

type dbEng struct{
	id int
	englishWord string
	russianWord string
	timeInit string
	timeCheck string
	examine int
	day int
	word string
	mistake int
	dayTime float64
	dayTime1 float64
	y int
}

var x int

func main() {
	for {
		str := new(dbEng)
		welcomeP()      // приветствие и инструкция
		scanX(&x)       // выбор варианта продолжения программы
		variant(x, str) // запуск одного из вариантов
	}
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
	fmt.Println("3 - выход из программы")
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

func scanY(str *dbEng) {
	var a string
	_, err := fmt.Scan(&a)
	str.y, err = strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}

// variant - запуск одного из вариантов
func variant(x int, str *dbEng){
	if x == 1 {
		checkWord(str)
	}
	if x == 2 {
		addWord(str)
		addWordDB(*str) // добавление слова в таблица
	}
	if x == 3 {
		os.Exit(0)
	}
}

// addWord - добавление слова
func addWord(str *dbEng) {
	fmt.Println("Введите слова по примеру:")
	fmt.Println("Word, Words")
	fmt.Println("Слово, Слова")
	scanWord(str)
	str.englishWord = str.word
	fmt.Println("str.englishWord: = ", str.englishWord)
	scanWord(str)
	str.russianWord = str.word
	fmt.Println("str.russianWord: = ", str.russianWord)

	// запись времени в переменные
	timeNow := time.Now()
	timeNow.String()
	str.timeInit = timeNow.Format("2006-01-02 15:04:05")
	str.timeCheck = timeNow.Format("2006-01-02 15:04:05")

	// количество повторений(inc)
	str.examine = 1

	// временной промежуток
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

	// количество повторных ошибок > 3 = откат на прошлый временной промежуток (day)
	str.mistake = 1

}

// scanWord - чтение слова из консоли
func scanWord(str *dbEng) {
	_, err := fmt.Scan(&str.word)
	if err != nil {
		fmt.Println(err)
	}
}

// addWorldDB = Добавляет новое слово в таблицу со стандартными значениями.
func addWordDB(str dbEng) {
	db, err := sql.Open("sqlite3", "english.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec(`
		insert into english (english_word, russian_word, time_init,
		time_check, examine, day, mistake) 
		values($1, $2, $3, $4, $5, $6, $7)`,
		str.englishWord, str.russianWord, str.timeInit, str.timeCheck,
		str.examine, str.day, str.mistake)
	if err != nil{
		panic(err)}
	fmt.Println(str.englishWord, str.russianWord, str.timeInit, str.timeCheck,
		str.examine, str.day)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

// checkWorld = Проверка слова
// Проверяет знание пользователя, изучил ли он слово за данный промежуток времени или нет.
func checkWord( str *dbEng) {
	db, err := sql.Open("sqlite3", "english.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from english")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []dbEng{}

	for rows.Next(){
		err := rows.Scan(&str.id, &str.englishWord, &str.russianWord, &str.timeInit, &str.timeCheck,
			&str.examine, &str.day, &str.mistake)
		if err != nil{
			fmt.Println(err)
			continue
		}
		products = append(products, *str)
	}
	for _, str := range products{
		fmt.Println(str.id, str.englishWord, str.russianWord, str.timeInit, str.timeCheck,
			str.examine, str.day, str.mistake)
		checkTime(&str)
	}
}

// checkTime - проверка слова на проверку
func checkTime(str *dbEng) {
	//форматирует время к нашему.
	timeNow := time.Now()
	timeNow.String()
	timeNow1 := timeNow.Format("2006-01-02 15:04:05")
	timeNow, err := time.Parse("2006-01-02 15:04:05", timeNow1)
	if err != nil {
		panic(err)
	}

	//перевод разницы во времени в секунды
	timeCheck, _ := time.Parse("2006-01-02 15:04:05", str.timeCheck)
	difference := timeNow.Sub(timeCheck)
	s, _ := time.ParseDuration(difference.String())
	str.dayTime1 = s.Seconds()

	//Здесь должна быть функция по проверки дня и равному времени для проверки.
	dayTime(str)
	timeCheckDay(str)
}

// временной промежуток
// day = 1 - 15 min
// day = 2 - 2 hours
// day = 3 - 1 day
// day = 4 - 3 day
// day = 5 - 7 day
// day = 6 - 21 day
// day = 7 - 50 day
// day = 8 - 150 day
// day = 9 - 365 day
func dayTime(str *dbEng) {
	//15m0s
	if str.day == 1{
		str.dayTime = 900 // 15min
	}
	if str.day == 2{
		str.dayTime = 7200 // 2h
	}
	if str.day == 3 {
		str.dayTime = 86400 // 1d
	}
	if str.day == 4 {
		str.dayTime = 259200 // 3d
	}
	if str.day == 5 {
		str.dayTime = 604800 // 7d
	}
	if str.day == 6 {
		str.dayTime = 1814400 // 21d
	}
	if str.day == 7 {
		str.dayTime = 4320000 // 50d
	}
	if str.day == 8 {
		str.dayTime = 12960000 // 150d
	}
	if str.day == 9 {
		str.dayTime = 31536000 // 365d
	}
}

// TimeCheckDay - вывод слова на проверку
func timeCheckDay(str *dbEng){
	a := str.dayTime1 - str.dayTime
	if a > 0 {
		str.examine = 2 // готово к проверку
		examineWord(str)
	} else {
		str.examine = 1 // не готово к проверке
		fmt.Println("Слово не готово для проверки")
	}
}

//Функция вывода слова и запроса перевода на него
func examineWord(str *dbEng){
	fmt.Println("Слово готово к проверке: ", str.englishWord)
	scanWord(str)
	fmt.Println("Ваш ответ: ", str.word, " Верный ответ: ", str.russianWord)
	fmt.Println("Если ответы совпадают, введите - 1")
	fmt.Println("Если ответы не совпадают, введите - 2")
	scanY(str)
}

//dateIn - ввод данных после экзамена
func dateIn(str *dbEng){

}