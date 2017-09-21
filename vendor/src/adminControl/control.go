package adminControl

import (
	dataId "data"
	"firebase"
	//	"fmt"
	"strconv"
	"structs"
	"telegram"
)

//admin commands
func CreateUser(id int64, dataFirebase *structs.Main) {
	if dataFirebase.Users == nil {
		dataFirebase.Users = map[string]*structs.User{}
	}
	dataFirebase.Users[strconv.FormatInt(id, 10)] = &structs.User{Id: id,
		Pay:      false,
		Speed:    0,
		Searches: map[string]*structs.Search{},
	}
	firebase.Update(dataFirebase)
	telegram.SendMessage("User created", dataId.ChatId)
}
func AddUser(data *structs.Main, id int64) {
	data.UsersId = append(data.UsersId, id)
	firebase.Update(data)
	telegram.SendMessage("User added", dataId.ChatId)
}
func CheckUser(data *structs.Main, id int64) bool {

	for _, val := range data.UsersId {
		if val == id {
			return true
		}
	}
	telegram.SendMessage("I don't know who you are. Please contact with my maker @taketas", id)
	return false
}
func CheckInSlice(data *structs.Main, id int64) bool {
	for _, val := range data.UsersId {
		if val == id {
			return true
		}
	}
	return false
}
func Pay(data *structs.Main, id int64) {
	if data.Users[strconv.FormatInt(id, 10)].Pay == true {
		data.Users[strconv.FormatInt(id, 10)].Pay = false
		firebase.Update(data)

	} else {
		data.Users[strconv.FormatInt(id, 10)].Pay = true
		firebase.Update(data)
	}

}
func Speed(speed int, data *structs.Main, id int64) {
	data.Users[strconv.FormatInt(id, 10)].Speed = speed
	firebase.Update(data)
}
func ShowInfo(data *structs.Main, id int64) {
	telegram.SendMessage("Pay: "+strconv.FormatBool(data.Users[strconv.FormatInt(id, 10)].Pay), dataId.ChatId)
	idStr := strconv.FormatInt(id, 10)
	speed := strconv.Itoa(data.Users[idStr].Speed)
	telegram.SendMessage("Speed: "+speed, dataId.ChatId)
	telegram.SendMessage("Searches: ", dataId.ChatId)
	for key, _ := range data.Users[idStr].Searches {
		telegram.SendMessage(key, dataId.ChatId)
	}
}
func ShowUsersIds(data *structs.Main) {
	telegram.SendMessage("Users: ", dataId.ChatId)
	for _, val := range data.UsersId {
		idStr := strconv.FormatInt(val, 10)
		telegram.SendMessage(idStr, dataId.ChatId)
	}
}

//func main() {
//	fmt.Println("vim-go")
//}
