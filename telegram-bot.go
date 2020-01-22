package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type dbEng struct{
	id int // порядковый номер в таблице
	englishWord string // английское слово
	russianWord string // русское слово - перевод
	timeInit string // дата создания записи
	timeCheck string // последняя проверка слова
	examine int // количество ошибок подряд (не может быть больше 2)
	day int // =(1-10) - период проверки слова
	word string // временное слово
	mistake int // количество ошибок (увеличивается всегда после теста)
	dayTime float64 // временная переменная для периода
	dayTime1 float64 // временная переменная для периода
	y int // временная переменая
	timeTemporal string // временная переменная для времени
	ready int // переменная готовности
}

var x int // надо снести, чтоб переписать на Y

func main() {
	bot, err := tgbotapi.NewBotAPI("1090414576:AAE-P5UDdzGngrzLjY7VXJ81P4R6LbGRFN0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		str := new(dbEng) // str - структура данных всей программы
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID
		variant(bot, &msg, update, str)
	}
}

func variant(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng  ) {
	if update.Message.Text == "1" {
		checkWord(bot, msg, update, str) // проверка слова
	}
}

func checkWord(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng) {
	msg.Text = "Проверка слов начата"
	bot.Send (msg)

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
		checkTime(bot, msg, update, &str) // готово ли слова для проверки
	}
}

func checkTime(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng){
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
	dayTime(str) // временной промежуток в секунды
	timeCheckDay(bot, msg, update, str) // готовность слова к проверке
}

func formatTime(str *dbEng) {
	timeNow := time.Now()
	timeNow.String()
	str.timeTemporal = timeNow.Format("2006-01-02 15:04:05")
}

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
	if str.day > 10 {
		str.dayTime = 157680000
	}
}

func timeCheckDay(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng){
	a := str.dayTime1 - str.dayTime
	if a > 0 {
		str.ready = 2 // готово к проверку
		examineWord(bot, msg, update, str) // вывод слова и запроса перевода на него
	} else {
		str.ready = 1 // не готово к проверке
		fmt.Println("Слово не готово для проверки")
	}
}

func examineWord(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng){
	msg.Text = "Слово готово к проверке: "+ str.englishWord
	scanWord(bot, msg, update, str)
	/*scanWord(str) // сканирование слова
	fmt.Println("Ваш ответ: ", str.word, " Верный ответ: ", str.russianWord)
	fmt.Println("Если ответы совпадают, введите - 1")
	fmt.Println("Если ответы не совпадают, введите - 2")
	scanY(str) // сканирование цифры
	examineEnd(str) // формирование данных для записи в базу данных
	dateIn(str) // запись в базу данных */
}

func scanWord(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update, str *dbEng){
	u := tgbotapi.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}