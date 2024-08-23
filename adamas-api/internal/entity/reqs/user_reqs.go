package reqs

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	Name        string `json:"name"`
	NickName    string `json:"nickname"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
