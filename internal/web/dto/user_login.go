package dto

type UserLoginDTO struct {
	Account  string `json:"account,required"`
	Password string `json:"password,required"`
}
