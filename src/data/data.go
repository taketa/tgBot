package data

const FirebaseUrl = "https://botmulti-18c0b.firebaseio.com/" //multi
const Used = "3000"
const ForParts = "7000"
const New = "1000"
const Phone = "9355"
const Shoes = "93427"

//This bot automates the process of searching for the new lots offered for sale on Ebay auction in the category "Cell phones & Smartphones". Access to these lots highly increases your chances of making a purchase faster and cheaper.
// If you want to use this bot press the "ID" button and send it to @taketas.
const Help = `
How to use:
1) Search format:
search+condition+minPrice+maxPrice+category
Note: if category not set - default value is "Cell Phones & Smartphones"
Available conditions:
1) "fp" - for parts or not working
2) "used"
3) "new"

Example:
1) iphone 7+fp+200+500
Searches for listings that have words "iphone" and "7" in title, "For parts or not working" in the category, with prices from $200 to $500.
2) "iphone 7"+used+200+500
Searches for listings that have the exact phrase "iphone" in the title, "Used" in the category, with prices from $200 to $500.
3) iphone -(T-mobile,cracked,icloud)+used+200+500
Searches for listings that have "iphone" in the title and do not have "T-mobile", "cracked", "icloud" in the title or description.

NOTE: If you request the "search request" that has already existed once againe it will be deleted. 

CATEGORIES REDUCTIONS
Watches = w
Smart Watches = sw
Video Endoscopes = ve
Cell Phone Accessories:
Accessory Bundles = abu
Armbands = aa
Audio Docs & Speakers = ads
Batteries = aba
Cables & Adapters = aca
Car Speakerphones = acs
Cases, Covers & Scins = accs
Cargers & Cradles = accr
Fceplates, Decals & Stickers = afds
Headsets = ah
Mounts & Holders = amm
Port Dust Covers = apdc
Screen Protectors = aspr
Straps & Charms = asch
Other Cell Phone Accessories = aother

`

const BotToken = "358478690:AAH1n-iye8x4l3Tabnh3tqnVqWJLJyJ-N6c" //bot for test
const ChatId = 328852780

//Telegram bot for automating the process of the searching for the new lots. (Beta)
//@taketasEbayBot - this bot automates the process of searching for the new lots offered for sale on Ebay auction with "Buy It Now" filter. Access to these lots highly increases your chances of making a purchase faster and cheaper.
//If you want to use this bot press the "ID" button and send it to @taketas.
//For more information contact @taketas.
//Available categories:
//Cell phones & Smartphones
//Smart Watches
//Video Endoscopes
//Cell Phone Accessories:
//Accessory Bundles
//Armbands
//Audio Docs & Speakers
//Batteries
//Cables & Adapters
//Car Speakerphones
//Cases, Covers & Scins
//Cargers & Cradles
//Fceplates, Decals & Stickers
//Headsets
//Mounts & Holders
//Port Dust Covers
//Screen Protectors
//Straps & Charms
//Other Cell Phone Accessories
//In feature the category list can be extended.
