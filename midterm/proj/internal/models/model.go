package models

import "regexp"

type Laptop struct {
	ID int `json:"id"`
	Name string `json:"name"`
	
	Year int `json:"year"`
	Display string `json:"display"`

	Memory string `json:"memory"`
	Storage string `json:"storage"`

	Condition string `json:"condition"`
	Description string `json:"description"`
	
	Phonenumber  int `json:"phonenumber"`
	
	Price string `json:"price"`
}

type Snowboard struct {
	ID int    `json:"id"`
	Name string `json:"name"`

	Size string `json:"size"`
	Condition string `json:"condition"`
	
	Description string `json:"description"`
	
	Phonenumber  int `json:"phonenumber"`
	Price string `json:"price"`
}

type Shirt struct {
	ID int    `json:"id"`
	Name string `json:"name"`

	Color string `json:"color"`
	Size string `json:"size"`
	Condition string `json:"condition"`

	Description string `json:"description"`
	Phonenumber  int `json:"phonenumber"`
	Price string `json:"price"`
}

type Toy struct {
	ID int    `json:"id"`
	Name string `json:"name"`

	Condition string `json:"condition"`
	Description string `json:"description"`
	
	Phonenumber  int `json:"phonenumber"`
	Price string `json:"price"`
}

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) IsEmailValid() bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(u.Email)
}


