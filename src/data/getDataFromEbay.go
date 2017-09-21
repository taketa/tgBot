package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"structs"
)

func GetNotes(url string) string {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer func() {
		resp.Body.Close()
	}()
	content := string(bytes)
	re := regexp.MustCompile("viSNotesCnt\">(.*)</span>")
	matches := re.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		return ""
	} else {
		return matches[0][1]
	}

	return matches[0][1]
}
func ConvertSearch(search string) (string, string, string, string, string) {
	retVal := strings.Split(search, "+")
	switch retVal[1] {
	case "fp":
		retVal[1] = "7000"
	case "used":
		retVal[1] = "3000"
	case "new":
		retVal[1] = "1000"
	}
	if len(retVal) == 4 {
		retVal = append(retVal, "9355")
	} else {

		switch retVal[4] {
		case "w": //Watches
			retVal[4] = "281"
		case "sw": //Smart Watches
			retVal[4] = "178893"
		case "ms": //Men's Shoes
			retVal[4] = "93427"
		case "msa": //Men's Shoes Athletic
			retVal[4] = "15709"
		case "msb": //Men's Shoes Boots
			retVal[4] = "11498"
		case "ln": //	Laptops & Netbooks
			retVal[4] = "175672"
		case "tr": //Tablets & eBook Readers
			retVal[4] = "171485"
		case "ve": //Video Endoscopes
			retVal[4] = "100001"
		//	Cell Phone Accessories
		case "abu": //Accessory Bundles
			retVal[4] = "58543"
		case "aa": //Armbands
			retVal[4] = ""
		case "ads": //Audio Docks & Speakers
			retVal[4] = "132297"
		case "aba": //Batteries
			retVal[4] = "20357"
		case "aca": //Cables & Adapters
			retVal[4] = "123422"
		case "acs": //Car Speakerphones
			retVal[4] = "88468"
		case "accs": //Cases, Covers & Skins
			retVal[4] = "20349"
		case "accr": //	Chargers & Cradles
			retVal[4] = "123417"
		case "afds": //Faceplates, Decals & Stickers
			retVal[4] = "20373"
		case "ah": //Headsets
			retVal[4] = "80077"
		case "amm": //Mounts & Holders
			retVal[4] = "35190"
		case "apdc": //Port Dust Covers
			retVal[4] = "179406"
		case "aspr": //Screen Protectors
			retVal[4] = "58540"
		case "asch": //Straps & Charms
			retVal[4] = "122962"
		case "aoter": //Other Cell Phone Accessories
			retVal[4] = "42425"
		}
	}
	return retVal[0], retVal[1], retVal[2], retVal[3], retVal[4]
}
func ItemValuesGet(single *structs.SingleItem) (*structs.ItemValues, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ItemValuesGet", r)
		}
	}()
	item := &structs.ItemValues{}
	item.Title = single.Item.Title
	item.StartTime = single.Item.StartTime
	item.ItemUrl = single.Item.ViewItemUrlForNaturalSearch
	item.SellerName = single.Item.Seller.UserId
	item.SellerFeeds = single.Item.Seller.FeedbackScore
	item.SellerPositiveFeedbackPercent = single.Item.Seller.PositiveFeedbackPercent
	//item.SellerRatingStar = single.Item.Seller.FeedbackRatingStar
	item.AutoPay = single.Item.AutoPay
	item.Description = single.Item.Description
	item.Returns = single.Item.ReturnPolicy.ReturnsAccepted
	item.BestOfferEnabled = single.Item.BestOfferEnabled
	item.BuyItNowAvailable = single.Item.BuyItNowAvailable
	if item.BuyItNowAvailable == true {
		item.BuyItNowPrice = single.Item.BuyItNowPrice.Value
	}
	item.CurrentPrice = single.Item.CurrentPrice.Value
	item.GalleryUrl = single.Item.GalleryUrl
	item.PictureUrl = single.Item.PictureUrl
	return item, nil
}

func GetId(search, condition, minPrice, maxPrice, category string) (string, error) {
	search = strings.Replace(search, " ", "%20", -1)
	urlGlobal := "http://svcs.ebay.com/services/search/FindingService/v1?OPERATION-NAME=findItemsAdvanced&SERVICE-VERSION=1.12.0&SECURITY-APPNAME=andreysa-test-PRD-44d8cb02c-2804a236&RESPONSE-DATA-FORMAT=JSON&paginationInput.entriesPerPage=1&keywords=" + search + "&itemFilter(0).name=Condition&itemFilter(0).value=" + condition + "&itemFilter(1).name=ListingType&itemFilter(1).value(0)=FixedPrice&itemFilter(1).value(1)=AuctionWithBIN&itemFilter(2).name=MinPrice&itemFilter(2).value=" + minPrice + "&itemFilter(3).name=MaxPrice&itemFilter(3).value=" + maxPrice + "&itemFilter(4).name=LocatedIn&itemFilter(4).value=US&sortOrder=StartTimeNewest&categoryId=" + category
	resp, err := http.Get(urlGlobal)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	////recover when index error
	//defer func() {
	//	if e := recover(); e != nil {
	//		fmt.Println("index error")
	//	}
	//}()
	var result structs.Response
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("###	Get data from the uri:###")
	fmt.Println("###	" + urlGlobal + "###")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in GetId", r)
		}
	}()

	itemId := result.FindItemsAdvancedResponse[0].SearchResult[0].Item[0].ItemId[0]

	return itemId, nil
}
func SingleItemGet(id string) (*structs.SingleItem, error) {
	singleItem := "http://open.api.ebay.com/shopping?callname=GetSingleItem&responseencoding=JSON&IncludeDetails=true&IncludeSelector=Details,Description,TextDescription,ItemSpecifics,Shipping&appid=andreysa-test-PRD-44d8cb02c-2804a236&version=515&ItemID=" + id
	resp, err := http.Get(singleItem)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result structs.SingleItem
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("###Get data from the uri:")
	fmt.Println(singleItem)
	return &result, nil

}
func CheckInDesc(text, search string) bool {
	if in, _ := regexp.MatchString(search, text); in == true {
		return true
	}
	return false
}
func ConvertSearchReq(str string) (mainSearch string, advOpt []string) {
	str = strings.ToLower(str)

	if test := strings.Index(str, "<"); test != -1 {
		mainSearch = str[:test-1]
		testEnd := strings.LastIndex(str, ">")
		adv := str[test+1 : testEnd]
		advOpt = strings.Split(adv, ",")
	} else {
		mainSearch = str
	}

	return
	//if test1 := strings.Index(str, "("); test1 != -1 {
	//	if test2 := strings.Index(str, ")"); test2 != -1 {
	//		oper := str[test1 : test2+1]
	//		return str[test2+2:], oper, 0
	//	} else {
	//		return "", "", 1
	//	}
	//} else {
	//	return str, "", 0
	//}
}
func xorArg(str string) (string, string) {
	ind1, ind2 := strings.Index(str, "("), strings.Index(str, ")")
	xor := str[ind1+1 : ind2]
	srch := str[ind2+2 : len(str)-1]
	xorArgs := strings.Split(xor, "&")
	var strRet string
	for _, val := range xorArgs {
		if strRet == "" {
			strRet += "(" + val
		} else {
			strRet += "|" + val
		}
	}
	strRet += ") " + srch
	return strRet, srch

}
func AdvancedSearch(mainSearch string, advOpt []string, desc string) bool {
	mainSearch = strings.ToLower(mainSearch)
	desc = strings.ToLower(desc)
	if advOpt != nil {
		for _, val := range advOpt {
			if strings.Index(val, "(") != -1 {

				strRet, srch := xorArg(val)
				if CheckInDesc(desc, srch) == true {
					if CheckInDesc(desc, strRet) == true {
						continue
					} else {
						fmt.Println("false 1")
						return false
					}

				} else {
					continue
				}
			}
			if CheckInDesc(desc, val) == true {
				fmt.Println("false 2", val, desc)
				return false
			} else {
				continue
			}
		}
		return true

	}

	return true

}
