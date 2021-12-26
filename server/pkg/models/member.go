package models

type Member struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	RegistrationDate string `json:"registrationDate"`
}
