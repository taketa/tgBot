package data

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertSearchReq(t *testing.T) {
	desc := "Samsung Galaxy S8+   Plus 64GB Orchid Gray (Verizon) SCREEN, Works"
	search := "\"galaxy s8\" plus <gg,tmobile,t-mobile,cracked,water damage,no power,not turn on,doesn't turn on,doesnt turn on,damaged,not working,broken,dead line pixels,dead pixel,cracks on,white spot,little spot,bad lcd,bad glass,small crack,big crack,cracked,dropped,phone is dead,not boot,burnt screen,burn-in,burn in,burned,not power on,screen crack,pink line,hardware issues,hardware issue,lcd issue,lines on,line on>"
	mainSearch, advOpt := ConvertSearchReq(search)
	fmt.Println("(((((", strings.Replace(mainSearch, "\"", "", -1))
	fmt.Println("(((((", strings.ToLower(desc))
	fmt.Println("!!!", strings.Index(strings.ToLower(desc), strings.Replace(mainSearch, "\"", "", -1)))
	fmt.Println(mainSearch, advOpt)
	for _, val := range advOpt {
		fmt.Println(val)
	}
	fmt.Println(AdvancedSearch(mainSearch, advOpt, desc))
	//fmt.Println(xorArg("\"(gg|dd) 64\""))
}
