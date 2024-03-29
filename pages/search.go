package pages

import (
	"github.com/sclevine/agouti"
)

//SearchPage element
type SearchPage struct {
	SearchInput *agouti.Selection
	SButton     *agouti.Selection
}

/*
Search page initialization
*/
func Search(page *agouti.Page) *SearchPage {
	search := SearchPage{
		SearchInput: page.FindByName("q"),
		SButton:     page.Find("asf")}
	return &search
}

/*
SearchWith text method
*/
func (s *SearchPage) SearchWith(text string) {
	s.SearchInput.SendKeys(text + "\uE007")
}
