package telegram

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func (r *Robot) SendMessage(chatId int64, text string) error {
	param := url.Values{
		"chat_id": {strconv.FormatInt(chatId, 10)},
		"text":    {text},
	}
	resp, err := http.PostForm(r.apiUrl+"sendMessage", param)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var respInfo struct {
		ResponseCommon
		Result Message
	}
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return err
	}

	if !respInfo.Ok {
		return respInfo.ResponseCommon
	}

	return nil
}
