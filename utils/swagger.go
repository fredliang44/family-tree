package utils

type Ping struct {
	Message string `json:"message" example:"Pong"`
	Code    int    `json:"code" example:"200"`
}

type GenCodeReq struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
	Phone    string `json:"phone" example:"17777777777"`
}

type GenResetCodeReq struct {
	Username string `json:"username" example:"username"`
}

type RegisterReq struct {
	Username   string `json:"username" example:"username"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

type ResetReq struct {
	Username   string `json:"username" example:"username"`
	Password   string `json:"password" example:"password"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

type StdResp struct {
	Message string `json:"message" example:"OK"`
	Code    int    `json:"code" example:"200"`
}

type VerifyResp struct {
	Message string `json:"message" example:"OK"`
	Code    int    `json:"code" example:"200"`
	Status  string `json:"status" example:"Verified"`
}

type ErrResp struct {
	Message string `json:"message" example:"DB Error Message"`
	Code    int    `json:"code" example:"400"`
}
