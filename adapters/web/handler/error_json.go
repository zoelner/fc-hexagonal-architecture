package handler

import "encoding/json"

func jsonError(msg string) []byte {
	Error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(Error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
