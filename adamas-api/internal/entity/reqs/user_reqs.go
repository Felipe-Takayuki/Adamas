package reqs

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
