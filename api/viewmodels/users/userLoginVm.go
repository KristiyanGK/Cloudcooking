package users

//UserLoginVM contains data from user login
type UserLoginVM struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
