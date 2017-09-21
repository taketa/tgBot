package main

import (
	"adminControl"
	dataP "data"
	"firebase"
	"fmt"
	"goroutines"
	"strconv"
	"strings"
	"structs"
	"telegram"
	"time"
	"userControl"
)

func main() {
	//send to telegram new lots
	var send = make(chan structs.Send, 1000000)
	//receive from telegram messages
	var receive = make(chan structs.Reseive, 1000000)
	fmt.Println("Hiiiiii!!!!")
	dataFirebase := firebase.Retrieve()
	fmt.Println("---dataFirebase", dataFirebase.Users["328852780"])
	if dataFirebase == nil {
		dataFirebase = &structs.Main{UsersId: []int64{}, Users: map[string]*structs.User{strconv.FormatInt(dataP.ChatId, 10): &structs.User{Id: dataP.ChatId,
			Pay:      false,
			Speed:    0,
			Searches: map[string]*structs.Search{},
		}}}
	} else {
		for key, data := range dataFirebase.Users {

			for search, _ := range data.Searches {
				keyConv, _ := strconv.ParseInt(key, 10, 64)
				dataFirebase.Users[strconv.FormatInt(keyConv, 10)].Searches[search].SearchChan = make(chan bool)
				go goroutines.GoSearch(search, dataFirebase, send, keyConv)
			}
		}
	}
	go telegram.BotStart(receive)
	go telegram.ListenSend(send)
	for {
		select {
		case get := <-receive:
			//@admin
			switch strings.ToLower(string(get.Text[0])) {
			case "+":
				if get.Id == dataP.ChatId {
					idInt64, _ := strconv.ParseInt(get.Text[1:], 10, 64)
					adminControl.AddUser(dataFirebase, idInt64)
					adminControl.CreateUser(idInt64, dataFirebase)
				} else {
					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question."+strconv.FormatInt(get.Id, 10), get.Id)
				}
				continue
			case "$":
				if get.Id == dataP.ChatId {
					idInt64, _ := strconv.ParseInt(get.Text[1:], 10, 64)
					if adminControl.CheckUser(dataFirebase, get.Id) == true {
						adminControl.Pay(dataFirebase, idInt64)
					} else {
						telegram.SendMessage("Try again taketa!", get.Id)

					}
				} else {

					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question.", get.Id)
				}
				continue
			case "=":
				if get.Id == dataP.ChatId {
					idInt64, _ := strconv.ParseInt(get.Text[1:], 10, 64)
					if adminControl.CheckUser(dataFirebase, idInt64) == true {
						adminControl.ShowInfo(dataFirebase, idInt64)
					} else {
						telegram.SendMessage("Try again taketa!", get.Id)

					}
				} else {

					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question.", get.Id)
				}
				continue
			case "@":
				if get.Id == dataP.ChatId {
					strSplit := strings.Split(string(get.Text), "@")
					speed, _ := strconv.Atoi(strSplit[2])
					idInt64, _ := strconv.ParseInt(strSplit[1], 10, 64)
					if adminControl.CheckUser(dataFirebase, idInt64) == true {
						adminControl.Speed(speed, dataFirebase, idInt64)
					} else {
						telegram.SendMessage("Try again taketa!", get.Id)

					}
				} else {

					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question.", get.Id)
				}
				continue
			case "?":
				if get.Id == dataP.ChatId {
					idInt64, _ := strconv.ParseInt(get.Text[1:], 10, 64)
					if adminControl.CheckInSlice(dataFirebase, idInt64) == true {
						telegram.SendMessage("YES", get.Id)
					} else {
						telegram.SendMessage("NO", get.Id)
					}
				} else {
					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question.", get.Id)
				}
				continue
			case "~":
				if get.Id == dataP.ChatId {
					adminControl.ShowUsersIds(dataFirebase)
				} else {
					telegram.SendMessage("Incorrect request. Try again or contact @taketas for asking question.", get.Id)
				}
				continue
			}
			//@user
			switch strings.ToLower(get.Text) {
			case "/start":
				if adminControl.CheckUser(dataFirebase, get.Id) == true {
					telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
					telegram.SendMessage(dataP.Help, get.Id)

				} else {
					telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
					telegram.SendMessage(dataP.Help, get.Id)
				}
			case "help":
				telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
				telegram.SendMessage(dataP.Help, get.Id)
			case "show":
				if adminControl.CheckUser(dataFirebase, get.Id) == true {
					telegram.SearchesKeyboard(dataFirebase, get.Id, get.Update)
				}
			case "clear":
				if adminControl.CheckUser(dataFirebase, get.Id) == true {
					userControl.ClearAllSearches(dataFirebase, get.Id)
					telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
				}
			case "clear all searches":
				if adminControl.CheckUser(dataFirebase, get.Id) == true {
					userControl.ClearAllSearches(dataFirebase, get.Id)
					telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
				}
			case "id":
				telegram.SendMessage(strconv.FormatInt(get.Id, 10), get.Id)
			case "back":
				telegram.StartKeyboard(dataFirebase, get.Id, get.Update)
			default:
				if adminControl.CheckUser(dataFirebase, get.Id) == true {
					for _, src := range strings.Split(get.Text, "\n") {
						test := userControl.Add(dataFirebase, src, send, get.Id)
						fmt.Println("Add:", src, test)
						time.Sleep(time.Second)
					}
				}

			}

		}
	}

}
