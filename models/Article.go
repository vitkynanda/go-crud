package models

import "time"

type Article struct {
	Id           	uint   		`json:"id"`
	Title        	string 		`json:"title"  binding:"required"`
	Content      	string 		`json:"content" binding:"required"`
	Category     	string 		`json:"category" binding:"required"`
	Created_date 	time.Time	`json:"created_date"`
	Updated_date	time.Time	`json:"updated_date"`
	Status			string		`json:"status" binding:"required"`
}