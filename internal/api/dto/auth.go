package dto

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserVerify struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserToken struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
