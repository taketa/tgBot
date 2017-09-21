package firebase

import (
	"data"
	"fmt"
	"github.com/melvinmt/firebase"
	//	"time"
	//	"reflect"

	//	toMap "github.com/fatih/structs"
	//	"reflect"
	"strconv"
	"strings"
	"structs"
)

//func Retrieve() map[int64]*structs.User {
//	url := data.FirebaseUrl
//
//	ref := firebase.NewReference(url).Export(false)
//	fmt.Println("@@@@@@@@@@@@@")
//	data := map[int64]*structs.User{}
//	if err := ref.Value(&data); err != nil {
//		panic(err)
//	}
//	for _, val := range data {
//		if val.Searches == nil {
//			val.Searches = map[string]*structs.Search{}
//		}
//	}
//	return data
//}
func Add(dataFirebase *structs.Main, search string, id int64) {
	if test := strings.Split(search, "+"); len(test) < 4 {
		fmt.Println("incorrect")
		//telegram.SendMessage("Incorrect search: "+search+". Try another.", id)
		return
	}
	for val, _ := range dataFirebase.Users[strconv.FormatInt(id, 10)].Searches {
		if search == val {
			fmt.Println("Has")
			//telegram.SendMessage("The search has already exists", id)
			return
		}
	}
	fmt.Println("addd")
	//telegram.SendMessage("Add search: "+search, id)
	//user := dataFirebaseR[id]
	//	if user.Searches == nil {
	//		user.Searches = []user.Search{}
	//
	//	}
	if dataFirebase.Users[strconv.FormatInt(id, 10)].Searches == nil {
		dataFirebase.Users[strconv.FormatInt(id, 10)].Searches = map[string]*structs.Search{}
		fmt.Println("in IF")
	}
	fmt.Println("heere")
	dataFirebase.Users[strconv.FormatInt(id, 10)].Searches[search] = &structs.Search{search, []string{}, make(chan bool)}
	fmt.Println(dataFirebase.Users[strconv.FormatInt(id, 10)].Searches[search])
	fmt.Println(dataFirebase)
	//	go goroutines.GoSearch(search, dataFirebase, send, id)
	Update(dataFirebase)

}
func CheckUser(dataFirebase map[int64]*structs.User, id int64) bool {

	if len(dataFirebase) > 0 {
		for key, _ := range dataFirebase {
			if key == id {
				return true
			}
		}
	}
	return false
}

//func CreateUser(id int64, dataFirebase *structs.Main) {
//	user := structs.User{}
//	user.Id = id
//	user.Pay = true
//	user.Speed = 0
//	user.Searches = map[string]*structs.Search{}
//	dataFirebase[id] = &user
//	fmt.Println("@@@@@@@", dataFirebase)
//	Update(dataFirebase)
//}
//func Clear(dataFirebase map[int64]*structs.User, id int64) {
//	fmt.Println("=======In the Clear========")
//	for _, i := range dataFirebase {
//		for _, val := range i.Searches {
//			fmt.Println("===sendToStopGoroutine===")
//			val.SearchChan <- true
//			<-val.SearchChan
//		}
//	}
//	//	buf := *dataFirebase
//	//	user := buf[id]
//	//	user.Searches = []user.Search{}
//	//	buf[id] = user
//	//	dataFirebase = &buf
//	dataFirebase[id].Searches = map[string]*structs.Search{}
//	Update(dataFirebase)
//
//}
func Show(data *structs.Main, id int64) {
	if data.Users[strconv.FormatInt(id, 10)].Searches == nil {
		fmt.Println("noooo")
		return
	}

	for search, _ := range data.Users[strconv.FormatInt(id, 10)].Searches {
		fmt.Println(search)
	}
	return
}

//====new func without firebase=====
func Retrieve() *structs.Main {
	url := data.FirebaseUrl
	ref := firebase.NewReference(url).Export(false)
	data := structs.Main{}
	if err := ref.Value(&data); err != nil {
		panic(err)
	}
	//for _, val := range data {
	//	if val.Searches == nil {
	//		val.Searches = map[string]*structs.Search{}
	//	}
	//}

	return &data
}
func Write(dataWrite *structs.User) error {
	url := data.FirebaseUrl
	ref := firebase.NewReference(url)

	if err := ref.Write(dataWrite); err != nil {
		return err
	}
	return nil
}
func Push(dataWrite string, where string) error {
	url := data.FirebaseUrl + where
	ref := firebase.NewReference(url)

	if err := ref.Push(dataWrite); err != nil {
		return err
	}
	return nil
}
func Update(dataWrite *structs.Main) error {
	url := data.FirebaseUrl
	ref := firebase.NewReference(url)

	if err := ref.Update(dataWrite); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//func MainToMainFire(data *structs.Main) *structs.MainFire {
//	dataRet := structs.MainFire{}
//
//	dataRet.Root = "taketa"
//
//	fmt.Println("data=====", data)
//	for key, _ := range data.Users {
//		fmt.Println(dataRet.Users[key].Id)
//		dataRet.Users[key].Id = key
//		dataRet.Users[key].Pay = data.Users[key].Pay
//		dataRet.Users[key].Speed = data.Users[key].Speed
//		for keyS, valS := range data.Users[key].Searches {
//			dataRet.Users[key].Searches[keyS].SearchReq = valS.SearchReq
//			dataRet.Users[key].Searches[keyS].SearchId = valS.SearchId
//
//		}
//	}
//	return &dataRet
//}

//func Transform(data map[int64]*structs.User) map[int64]interface{} {
//	var dataRet map[int64]interface{}
//	for key, val := range data {
//		fmt.Println("kkkkeeeeeeeeeeeeeeey", key)
//		dataRet[key]=map[string]interface{}
//		dataRet[key] = toMap.Map(val)
//		fmt.Println("obrabotan")
//	}
//
//	return dataRet
//}
func SearchesTransform(data map[string]*structs.Search) map[string]structs.Search {
	dataRet := map[string]structs.Search{}
	for key, val := range data {
		dataRet[key] = *val
	}
	return dataRet
}

//func UpdateSearches(id int64, dataWrite *map[string][]string) error {
//	url := data.FirebaseUrl + string(id) + "/Searches/"
//	ref := firebase.NewReference(url)
//
//	if err := ref.Update(dataWrite); err != nil {
//		return err
//	}
//	return nil
//}
func AppendUnic(dataBase *structs.Main, id int64, search string, what string) {
	fmt.Println("=====AppendUnic")

	dataBase.Users[strconv.FormatInt(id, 10)].Searches[search].SearchId = append(dataBase.Users[strconv.FormatInt(id, 10)].Searches[search].SearchId, what)
	Update(dataBase)
}

//func CheckForUnic(arr []string, unic string) bool {
//	for _, val := range arr {
//		if val == unic {
//			return false
//		}
//	}
//	return true
//}
func IdUnic(dataS []string, id string) bool {
	for _, val := range dataS {
		if val == id {
			return false
		}
	}
	return true
}

//===old func===
//func Update(url string, search interface{}) error {
//	url = strings.Replace(url, " ", "%20", -1)
//
//	url = strings.Replace(url, "+", "%2B", -1)
//	url = strings.Replace(url, "\"", "%22", -1)
//	ref := firebase.NewReference(url)
//
//	if err := ref.Update(search); err != nil {
//		return err
//	}
//	return nil
//}
//
//func Push(url string, search interface{}) error {
//	url = strings.Replace(url, " ", "%20", -1)
//
//	url = strings.Replace(url, "+", "%2B", -1)
//	url = strings.Replace(url, "\"", "%22", -1)
//	ref := firebase.NewReference(url)
//
//	if err := ref.Push(search); err != nil {
//		return err
//	}
//	return nil
//}
//
func Delete(url string) error {
	url = strings.Replace(url, " ", "%20", -1)

	url = strings.Replace(url, "+", "%2B", -1)
	url = strings.Replace(url, "\"", "%22", -1)

	ref := firebase.NewReference(url)

	if err := ref.Delete(); err != nil {
		return err
	}
	return nil
}

//func GetSearches() []string {
//	url := data.FirebaseUrl + "searches"
//
//	ref := firebase.NewReference(url).Export(false)
//	var ids map[string]interface{}
//	if err := ref.Value(&ids); err != nil {
//		panic(err)
//	}
//	searches := []string{}
//	for k, _ := range ids {
//		searches = append(searches, k)
//	}
//
//	return searches
//}

//func main() {
//	test := Retrieve()
//	//	fmt.Println(test[123].Searches)
//	//	fmt.Println(test[123].Searches == nil)
//	//	fmt.Println(CheckUser(test, 124))
//	//	CreateUser(222, test)
//	//	Clear(test, 123)
//	fmt.Println("test: ", &test)
//	//	Show(test, 123)
//	//	fmt.Println(MainToMainFire(test))
//	//	Show(test, 123)
//	//	Transform(test)
//	Add(test, "iphone 6s+fp+10+50", 123)
//}
