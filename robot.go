package telegram

import (
	"fmt"
)

const ApiUrl = "https://api.telegram.org/bot"

type Robot struct {
	apiUrl string
}

func New(token string) *Robot {
	return &Robot{apiUrl: ApiUrl + token + "/"}
}

type ResponseCommon struct {
	Ok          bool
	Description string
	ErrorCode   int `json:"error_code"`
}

func (r ResponseCommon) Error() string {
	return fmt.Sprintf("[%d]%s", r.ErrorCode, r.Description)
}
