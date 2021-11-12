package models

import "regexp"

type Laptop struct {
	ID int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	
	Year int `json:"year" db:"year"`
	Display string `json:"display" db:"display"`

	Memory string `json:"memory" db:"memory"`
	Storage string `json:"storage" db:"storage"`

	Condition string `json:"condition" db:"condition"`
	Description string `json:"description" db:"description"`
	
	Phonenumber  string `json:"phonenumber" db:"phonenumber"`
	
	Price string `json:"price" db:"price"`
}

type Snowboard struct {
	ID int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Size string `json:"size" db:"size"`
	Condition string `json:"condition" db:"condition"`

	Description string `json:"description" db:"description"`

	Phonenumber  string `json:"phonenumber" db:"phonenumber"`
	Price string `json:"price" db:"price"`
}

type Shirt struct {
	ID int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Color string `json:"color" db:"color"`
	Size string `json:"size" db:"size"`
	Condition string `json:"condition" db:"condition"`

	Description string `json:"description" db:"description"`
	Phonenumber  string `json:"phonenumber" db:"phonenumber"`
	Price string `json:"price" db:"price"`
}

type Toy struct {
	ID int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Condition string `json:"condition" db:"condition"`
	Description string `json:"description" db:"description"`

	Phonenumber  string `json:"phonenumber" db:"phonenumber"`
	Price string `json:"price" db:"price"`
}

type User struct {
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *User) IsEmailValid() bool {
   emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
   return emailRegex.MatchString(u.Email)
}


