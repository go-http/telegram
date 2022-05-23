package telegram

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (r *Robot) GetUpdates() ([]Update, error) {
	param := url.Values{
		"timeout": {"20"},
	}
	resp, err := http.PostForm(r.apiUrl+"getUpdates?offset=-3&timeout=200", param)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respInfo struct {
		ResponseCommon
		Result []Update
	}
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	if !respInfo.Ok {
		return nil, respInfo.ResponseCommon
	}

	return respInfo.Result, nil
}
