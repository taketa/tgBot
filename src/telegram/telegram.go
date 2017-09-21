package telegram

import (
	//	"adminControl"
	"data"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"structs"
)

var bot, Err = tgbotapi.NewBotAPI(data.BotToken)

func SearchesKeyboard(dataB *structs.Main, id int64, botUp tgbotapi.Update) {
	msg := tgbotapi.NewMessage(botUp.Message.Chat.ID, botUp.Message.Text)
	var buttons []tgbotapi.KeyboardButton
	var butRow [][]tgbotapi.KeyboardButton
	for key, _ := range dataB.Users[strconv.FormatInt(id, 10)].Searches {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(key))
		fmt.Println("*******************", buttons)

	}

	for _, btn := range buttons {
		butRow = append(butRow, tgbotapi.NewKeyboardButtonRow(btn))
	}
	//buttons = append(buttons, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("clear")))
	butRow = append(butRow, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("clear all searches")))
	butRow = append(butRow, tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("back")))
	keyboard := tgbotapi.NewReplyKeyboard(butRow...)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)

}
func StartKeyboard(dataB *structs.Main, id int64, botUp tgbotapi.Update) {
	msg := tgbotapi.NewMessage(botUp.Message.Chat.ID, botUp.Message.Text)
	var numericKeyboard tgbotapi.ReplyKeyboardMarkup
	if id == data.ChatId {

		numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("show"),
				tgbotapi.NewKeyboardButton("clear"),
				tgbotapi.NewKeyboardButton("help"),
				tgbotapi.NewKeyboardButton("~"),
			),
			//tgbotapi.NewKeyboardButtonRow(
			//tgbotapi.NewKeyboardButton(""),
			//tgbotapi.NewKeyboardButton(""),
			//),
		)
	} else {
		numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("show"),
				tgbotapi.NewKeyboardButton("clear"),
				tgbotapi.NewKeyboardButton("id"),
				tgbotapi.NewKeyboardButton("help"),
			),
		)
	}
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}
func BotStart(channel chan structs.Reseive) {
	if Err != nil {
		log.Panic(Err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//		fmt.Println("id", update.Message.Chat.ID)
		//		keyboard := startKeyboard(update.Message.Chat.ID)
		//
		//		msg.ReplyMarkup = keyboard
		//
		//		fmt.Println("@@@@@@@@@@@MSG2@@@@", msg)
		//		//bot.Send(msg)
		//
		channel <- structs.Reseive{Text: update.Message.Text, Id: update.Message.Chat.ID, Update: update}
		//channel <- map[string]interface{}{"text": update.Message.Text, "id": update.Message.Chat.ID}
	}
}

func SendSearches(searches []string, id int64) {
	if Err != nil {
		log.Panic(Err)
	}
	for i, search := range searches {
		msg := tgbotapi.NewMessage(int64(id), string(i)+" "+search)
		bot.Send(msg)
	}
}
func SendMessage(message string, id int64) {
	if Err != nil {
		log.Panic(Err)
	}

	msg := tgbotapi.NewMessage(id, message)

	bot.Send(msg)
}
func ListenSend(sendCh chan structs.Send) {
	for {
		receive := <-sendCh
		SendMessage(receive.Text, receive.Id)
	}
}
