package utils

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"

	"go.trulyao.dev/lito/pkg/errors"
)

var allowedChars = []rune("abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ123456789!=_-")

type Response struct {
	OK      bool        `json:"ok"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Err     error       `json:"-"`
}

func GenerateAPIKey() string {
	return GenerateRandomString(32)
}

func GenerateRandomString(length int) string {
	str := make([]byte, length)

	n, err := io.ReadAtLeast(rand.Reader, str, length)
	if n != length || err != nil {
		return ""
	}

	for i := 0; i < length; i++ {
		str[i] = byte(allowedChars[int(str[i])%len(allowedChars)])
	}

	return string(str)
}

func SendJSON(w http.ResponseWriter, res *Response) {
	if res == nil {
		res = &Response{
			OK:   true,
			Code: http.StatusOK,
		}
	}

	if res.Code == 0 {
		res.Code = http.StatusOK
	}

	if res.OK && res.Code >= 400 {
		res.OK = false
	}

	if !res.OK && res.Code < 400 {
		res.OK = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		SendErrorResponse(w, err)
		return
	}

	return
}

func SendOK(w http.ResponseWriter, data interface{}) {
	SendJSON(w, &Response{
		OK:   true,
		Code: http.StatusOK,
		Data: data,
	})
}

func SendErrorResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	if err != nil {
		if cErr, ok := err.(*errors.APIError); ok {
			message = cErr.Error()
			code = cErr.Code()
		}
	}

	SendJSON(w, &Response{
		OK:    false,
		Code:  code,
		Error: message,
		Err:   err,
	})
}
