package models


type User struct {
ID int32 `json:"id"`
Name string `json:"name" validate:"required"`
DOB string `json:"dob" validate:"required,datetime=2006-01-02"`
Age int `json:"age,omitempty"`
}