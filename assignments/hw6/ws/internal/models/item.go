package models

type Item struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`

	Price string `json:"price"`
	Condition string `json:"condition"`
	Description    string `json:"description"`
	
	Author  string `json:"author"`
	Phonenumber  int `json:"Phonenumber"`
}
