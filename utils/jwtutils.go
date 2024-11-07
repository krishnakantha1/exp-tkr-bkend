package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krishnakantha1/expenseTrackerBackend/types"
)

/*
function used to get the jwt from a http request

params:

	r *http.Request

return:

	string, error
*/
func GetJWTFromHeader(r *http.Request) (string, error) {
	val := r.Header.Get("Authorization")

	if len(val) == 0 {
		return "", errors.New("authorization headder not found")
	}

	//val should have the format: Bearer <TOKEN>
	parts := strings.Split(val, "Bearer ")

	if len(parts) != 2 {
		return "", errors.New("could not parse Bearer info")
	}

	return parts[1], nil
}

/*
given a JWT string function returns the parsed JWTUser type. If bad `jwtstr“ provided, returns error.

params:

	jwtstr string

returns:

	*types.JWTUser, error
*/
func GetUserFromJWT(jwtstr string) (*types.JWTUser, error) {
	user := &types.JWTUser{}

	key, err := GetEnv("PRIVATEKEY")
	if err != nil {
		return nil, err
	}

	_, err = jwt.ParseWithClaims(jwtstr, user, func(_ *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

/*
given a JWTUser, returns a JWT string consturcted with the claims obtained from the JWTUser.GetClaims method call.

params:

	user *types.JWTUser

return:

	string
*/
func GetJWTFromUser(user *types.JWTUser) (string, error) {
	//default expiration is 7 days
	user.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour))

	iss, err := GetEnv("ISSUER")
	if err != nil {
		return "", err
	}
	user.Issuer = iss

	key, err := GetEnv("PRIVATEKEY")
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, user)

	return t.SignedString([]byte(key))
}
