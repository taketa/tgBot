package goroutines

import (
	"data"
	"firebase"
	"fmt"
	"instantView"
	"strconv"
	"structs"
	"telegram"
	"time"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}
func GoSearch(search string, dataB *structs.Main, send chan structs.Send, sendId int64) {
	var counter int64
	counter = 0
	if dataB.Users[strconv.FormatInt(sendId, 10)].Pay == false {
		telegram.SendMessage("Please donate or contact @taketas.", sendId)
	}

	fmt.Println("SSSSSSSSSSSSSSSSSS", search)
	for {
		select {
		case <-dataB.Users[strconv.FormatInt(sendId, 10)].Searches[search].SearchChan:
			fmt.Println("=====receive======")
			dataB.Users[strconv.FormatInt(sendId, 10)].Searches[search].SearchChan <- true
			fmt.Println("stop search: ", search)
			return
		default:
			if dataB.Users[strconv.FormatInt(sendId, 10)].Pay == true {
				//time.Sleep(time.Second)
				sr, sr1, sr2, sr3, sr4 := data.ConvertSearch(search)
				mainSearch, _ := data.ConvertSearchReq(sr)
				id, _ := data.GetId(mainSearch, sr1, sr2, sr3, sr4)
				fmt.Println("@@@@@@@ID@@@@===", id)
				if id == "" {
					//telegram.SendMessage("There is no items for the search: "+search, sendId)
					//delete(dataB.Users[strconv.FormatInt(sendId, 10)].Searches, search)
					//firebase.Update(dataB)
					//return

					continue
				}
				if firebase.IdUnic(dataB.Users[strconv.FormatInt(sendId, 10)].Searches[search].SearchId, id) == true {
					singleItem, err := data.SingleItemGet(id)
					for err != nil {
						singleItem, err = data.SingleItemGet(id)
					}
					itemValues, err := data.ItemValuesGet(singleItem)
					fmt.Println("====VALUES----", "search=", search, &send, itemValues)
					for err != nil {

						itemValues, err = data.ItemValuesGet(singleItem)
					}
					firebase.AppendUnic(dataB, sendId, search, id)
					//		firebase.UpdateSearches(sendId, &dataB)
					var price string
					//	defer func() {
					//		if e := recover(); e != nil {
					//			go GoSearch(search, dataB, send, sendId)
					//			fmt.Println("@%%%%%%%%%%%%%%%%")
					//			//send<-map[string]interface{}{"text":"has no items for "}
					//		}

					//	}()
					fmt.Println("=======!!!!====ItemValues", "SEARCH=", search, "SEND=", &send, itemValues)
					if itemValues.BuyItNowAvailable == true {
						price = FloatToString(itemValues.BuyItNowPrice)
					} else {
						price = FloatToString(itemValues.CurrentPrice)
					}
					mainSearch, advOpt := data.ConvertSearchReq(search)
					if data.AdvancedSearch(mainSearch, advOpt, itemValues.Title+itemValues.Description+data.GetNotes(itemValues.ItemUrl)) == true {
						if counter == 1 {
							time.Sleep(time.Second * time.Duration(dataB.Users[strconv.FormatInt(sendId, 10)].Speed))
						}

						//!!!time when lot is reseived
						//"Start time: " + itemValues.StartTime + "\nReceive time: " + time.Now().String() + "\n" +
						commitToBuy := func() string {

							if itemValues.AutoPay == true {
								return "no"
							} else {
								return "yes"
							}
						}()
						bestOffer := func() string {
							if itemValues.BestOfferEnabled == true {
								return "yes"
							} else {
								return "no"
							}
						}()
						var pictures string
						for _, pic := range itemValues.PictureUrl {
							pictures += "<img src=\"" + pic + "\"/><br>"
						}
						forSend := "$" + price + "<br>" + "Rate: " + FloatToString(itemValues.SellerFeeds) + "/" + FloatToString(itemValues.SellerPositiveFeedbackPercent) + "<br>CommitToBuy: " + commitToBuy + "<br>" + "BestOffer: " + bestOffer + "<br>Search: " + search + "<br>" + pictures
						//func() string {
						//      for _, ur := range itemValues.PictureUrl {
						//              telegram.SendMessage(ur, sendId)
						//      }
						//      return "URL:\n"
						//}()

						send <- structs.Send{"Title:" + itemValues.Title + "\n$" + price + "\n" + "Rate: " + FloatToString(itemValues.SellerFeeds) + "/" + FloatToString(itemValues.SellerPositiveFeedbackPercent) + "\nCommitToBuy: " + commitToBuy + "\n" + "BestOffer: " + bestOffer + "\nSearch: " + search + "\n" + instantView.Instant(itemValues.Title, forSend+itemValues.Description+"<p>NOTES:</p>"+"<p>"+itemValues.Notes+"</p>"+"<p><a href=\""+itemValues.ItemUrl+"\">"+itemValues.ItemUrl+"</a></p>") + "\n\n" + itemValues.ItemUrl, sendId}

						counter = 1
					}

				}
			}
		}
	}
}

//func main() {
//	fmt.Println("vim-go")
//	fmt.Println(IdUnic("searches/3", "11"))
//
//}
