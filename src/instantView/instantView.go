package instantView

import (
	"log"

	//	"fmt"
	telegraph "github.com/toby3d/go-telegraph"
)

// Example content. Be sure to wrap every media in a <figure> tag, okay?
// Be easy, bro.

func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Instant(title, text string) string {
	var data string
	//	for _, img := range images {
	//		data += "<figure><img src=\"" + img + "\"/></figure>"
	//	}
	data += "<p>" + text + "</p>"
	// Create new Telegraph account. Author name/link can be epmty.
	// So secure. Much anonymously. Wow.
	newAccount := &telegraph.Account{
		ShortName:  "taketas", // required
		AuthorName: "taketa pada",
		//AuthorURL:  "https://t.me/toby3d",
	}
	acc, err := telegraph.CreateAccount(newAccount)
	checkError(err)

	// Boom!.. And your text will be understandable for Telegraph. MAGIC.
	content, err := telegraph.ContentFormat(data)
	checkError(err)

	newPage := &telegraph.Page{
		Title:   title,
		Content: content,

		// Not necessarily, but, hey, it's just an example.
		AuthorName: acc.AuthorName,
	}

	page, err := acc.CreatePage(newPage, false)
	checkError(err)

	log.Println("Kaboom! Page created, look what happened:", page.URL)
	return page.URL
}
