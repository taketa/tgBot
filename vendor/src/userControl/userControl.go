package userControl

import (
	"firebase"
	"fmt"
	"goroutines"
	"strconv"
	"strings"
	"structs"
	"telegram"
)

func Show(data *structs.Main, id int64) {
	if data.Users[strconv.FormatInt(id, 10)].Searches == nil {

		telegram.SendMessage("No searches", id)
		return
	}
	telegram.SendMessage("Searches:", id)
	for search, _ := range data.Users[strconv.FormatInt(id, 10)].Searches {
		telegram.SendMessage(search, id)
	}
	return
}
func ClearAllSearches(dataFirebase *structs.Main, id int64) {
	for srch, val := range dataFirebase.Users[strconv.FormatInt(id, 10)].Searches {
		fmt.Println("===sendToStopGoroutine===")
		val.SearchChan <- true
		<-val.SearchChan
		delete(dataFirebase.Users[strconv.FormatInt(id, 10)].Searches, srch)

	}

	firebase.Update(dataFirebase)
	telegram.SendMessage("Clear all searches", id)
}
func Add(dataFirebase *structs.Main, search string, send chan structs.Send, id int64) bool {

	fmt.Println("SEARCHHHHH!!!!!", search)
	if test := strings.Split(search, "+"); len(test) < 4 {

		fmt.Println("SEARCHHHHH!!!!!", test)
		telegram.SendMessage("Incorrect search: "+search+". Try another.", id)
		return false
	}
	for val, srch := range dataFirebase.Users[strconv.FormatInt(id, 10)].Searches {
		if search == val {
			fmt.Println("===sendToStopGoroutine===")
			srch.SearchChan <- true
			<-srch.SearchChan

			delete(dataFirebase.Users[strconv.FormatInt(id, 10)].Searches, val)
			telegram.SendMessage("Remove search: "+val, id)
			firebase.Update(dataFirebase)
			return true
		}
	}
	telegram.SendMessage("Add search: "+search, id)
	if dataFirebase.Users[strconv.FormatInt(id, 10)].Searches == nil {
		dataFirebase.Users[strconv.FormatInt(id, 10)].Searches = map[string]*structs.Search{}
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!HHHHHHHHHHHHHHHHH!!!!!!!")
	dataFirebase.Users[strconv.FormatInt(id, 10)].Searches[search] = &structs.Search{search, []string{}, make(chan bool)}

	go goroutines.GoSearch(search, dataFirebase, send, id)
	firebase.Update(dataFirebase)
	return false
}
