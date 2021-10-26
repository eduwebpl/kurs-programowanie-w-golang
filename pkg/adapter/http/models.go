package http

type CreateUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string `json:"accessToken"`
}

type UserInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
