package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/ilhamnyto/twinyto/config"
)

var (
	secret_key = config.GetString(config.SECRET_KEY)
	token_expired_time = 10 * 60 * time.Second
)

func GenerateToken(Userid int) (string, error) {
	payload := Payload{
		UserId: Userid,
		Expired: time.Now().Add(token_expired_time),
	}

	claims := jwt.MapClaims{
		"payload":payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret_key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method of %v", t.Header["alg"])
		}

		return []byte(secret_key), nil
	} )

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloadInterface := claims["payload"]

		payload := Payload{}

		payloadByte, err := json.Marshal(payloadInterface)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(payloadByte, payload)

		if err != nil {
			return nil, err
		}

		return &payload, nil
	} else {
		return nil, errors.New("Invalid token.")
	}
}