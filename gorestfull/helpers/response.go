package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func Response(w http.ResponseWriter, data interface{}, code int, isErr bool) {

	if isErr {
		json.NewEncoder(w).Encode(&response{
			Status:      getStatus(code),
			IsError:     isErr,
			Description: data,
		})
		return
	} else {
		json.NewEncoder(w).Encode(&response{
			Status:  getStatus(code),
			IsError: isErr,
			Data:    data,
		})
		return
	}
}

func getStatus(status int) string {
	var desc string

	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
