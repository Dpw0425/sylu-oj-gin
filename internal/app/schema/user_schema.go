package schema

type UserRegister struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}
