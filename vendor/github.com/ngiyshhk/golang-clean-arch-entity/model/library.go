package model

// Library is
type Library struct {
	Libid      string `json:"libid"`
	Libkey     string `json:"libkey"`
	Category   string `json:"category"`
	City       string `json:"city"`
	Short      string `json:"short"`
	Tel        string `json:"tel"`
	Pref       string `json:"pref"`
	Faid       string `json:"faid"`
	Geocode    string `json:"geocode"`
	Systemid   string `json:"systemid"`
	Address    string `json:"address"`
	Post       string `json:"post"`
	URLPc      string `json:"url_pc"`
	Systemname string `json:"systemname"`
	Isil       string `json:"isil"`
	Formal     string `json:"formal"`
}

// Libraries is
type Libraries = []*Library
