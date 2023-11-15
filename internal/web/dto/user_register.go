package dto

type UserRegisterDTO struct {
	Account       string `json:"account,required"`
	Password      string `json:"password,required"`
	CheckPassword string `json:"check_password,required"`
	PlanetCode    string `json:"planet_code,required"`
}
