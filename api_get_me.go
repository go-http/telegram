package telegram

import (
	"encoding/json"
	"net/http"
)

func (r *Robot) GetMe() (*User, error) {
	resp, err := http.Get(r.apiUrl + "getMe")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respInfo struct {
		ResponseCommon
		Result User
	}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	if respInfo.Ok {
		return &respInfo.Result, nil
	}

	return nil, respInfo
}
