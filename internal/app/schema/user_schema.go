package schema

type UserRegister struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseMenu struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}
