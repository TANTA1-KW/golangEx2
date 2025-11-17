package model

type Users struct{
	Email string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineID string `json:"line_id"`
	Phone string `json:"phone" validate:"required,numeric,len=10"`
	Business string `json:"business" validate:"required,bustype"`
	Website string `json:"website" validate:"required,min=2,max=30,url"`
}

