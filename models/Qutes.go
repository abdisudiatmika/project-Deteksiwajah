package models

//Qutes menyimpan QUtes
type Qutes struct {
	Count   int `json:"count"`
	Results []struct {
		QuoteAuthor string `json:"quoteAuthor"`
		QuoteText   string `json:"quoteText"`
	} `json:"results"`
}
