package types

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
