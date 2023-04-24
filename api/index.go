package handler

import (
	"encoding/json"
	"fmt"
	"github.com/acheong08/OpenAIAuth/auth"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	auth := auth.NewAuthenticator(r.Header.Get("email"), r.Header.Get("password"), r.Header.Get("proxy"))
	err := auth.Begin()
	ret := map[string]any{}

	if err.Error != nil {
		ret["message"] = err.Details
		responseJson(w, ret)
		return
	}
	token, err := auth.GetAccessToken()
	if err.Error != nil {
		ret["message"] = err.Details
		responseJson(w, ret)
		return
	}

	ret["access_token"] = token
	responseJson(w, ret)
}

func responseJson(w http.ResponseWriter, data map[string]interface{}) {
	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, `{"message": "server error"}`)
		return
	}
	fmt.Fprintf(w, "%s", marshal)
}
