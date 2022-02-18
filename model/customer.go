package model

type Customer struct {
	APIKEY    string `json:"api_key"`
	COMPANY   string `json:"company"`
	FIRSTNAME string `json:"first_name"`
	LASTNAME  string `json:"last_name"`
}
