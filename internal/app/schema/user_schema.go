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

type UserInfo struct {
	Identity string         `json:"identity"`
	Info     AccountInfo    `json:"info"`
	Menu     []ResponseMenu `json:"menu"`
	Students []string       `json:"students"`
}

type AccountInfo struct {
	Username  string `json:"username"`
	StartTime string `json:"startTime"`
	Submit    int64  `json:"submit"`
	Accept    int64  `json:"Accept"`
}
