package structs

import "github.com/go-telegram-bot-api/telegram-bot-api"

//user struct
type Main struct {
	UsersId []int64
	Users   map[string]*User
}
type User struct {
	Id       int64
	Pay      bool
	Speed    int
	Searches map[string]*Search
}
type Search struct {
	SearchReq  string
	SearchId   []string
	SearchChan chan bool `json:"-"`
}

//forFirebaseWrite

type MainFire struct {
	Root  string
	Users map[int64]UserFire
}
type UserFire struct {
	Id       int64
	Pay      bool
	Speed    int
	Searches map[string]SearchFire
}
type SearchFire struct {
	SearchReq string
	SearchId  []string
}

//receive from telegram messages
type Reseive struct {
	Text   string
	Id     int64
	Update tgbotapi.Update
}
type Send struct {
	Text string
	Id   int64
}

//structs for getting ID
type Response struct {
	FindItemsAdvancedResponse []*Item
}
type Item struct {
	SearchResult []*ItemAloneOne
}
type ItemAloneOne struct {
	Item []*ItemAlone
}
type ItemAlone struct {
	ItemId []string
}

//structs for SingleItem
type SingleItem struct {
	Item *Values
}
type Values struct {
	Title                       string
	ViewItemUrlForNaturalSearch string
	StartTime                   string
	Seller                      *Seller
	AutoPay                     bool
	Description                 string
	ReturnPolicy                *ReturnPolicy
	BestOfferEnabled            bool
	BuyItNowAvailable           bool
	BuyItNowPrice               *PriceBIN
	CurrentPrice                *CurrentPrice
	ItemId                      string
	GalleryUrl                  string
	PictureUrl                  []string
	Notes                       string
}
type CurrentPrice struct {
	Value float64
}
type Seller struct {
	UserId                  string
	FeedbackScore           float64
	PositiveFeedbackPercent float64
	FedbackRatingStar       string
}
type ReturnPolicy struct {
	ReturnsAccepted string
}
type PriceBIN struct {
	Value float64
}
type ItemValues struct {
	Title                         string
	ItemUrl                       string
	StartTime                     string
	SellerName                    string
	SellerFeeds                   float64
	SellerPositiveFeedbackPercent float64
	SellerRatingStar              string
	AutoPay                       bool
	Description                   string
	Returns                       string
	BestOfferEnabled              bool
	BuyItNowAvailable             bool
	BuyItNowPrice                 float64
	CurrentPrice                  float64
	GalleryUrl                    string
	PictureUrl                    []string
	Notes                         string
}

//for firebase
type FirebaseStruct struct {
	Search map[string][]string
}
type RetrFirebase struct {
	Ret []string
}
