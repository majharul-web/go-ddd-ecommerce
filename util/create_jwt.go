package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub          int    `json:"sub"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	IsShopOOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerBase64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadBase64 := base64UrlEncode(byteArrData)

	byteArrSecret := []byte(secret)

	message := headerBase64 + "." + payloadBase64
	byteArrMessage := []byte(message)

	hash := hmac.New(sha256.New, byteArrSecret)
	hash.Write(byteArrMessage)
	signature := hash.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)

	jwt := strings.Join([]string{headerBase64, payloadBase64, signatureBase64}, ".")
	return jwt, nil

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
