package jwt

import "net/http"

// type Claims struct {
// 	jwt.StandardClaims

// }

// // AccessToken creates a new access token and returns it in both JWT and
// // signed format, along with any error
// func AccessToken() (*jwt.Token, string, error) {

// }

type errorHandler func(w http.ResponseWriter, r *http.Request, err string)

// TokenExtractor is a function that takes a request as input and returns
// either a token or an error
type TokenExtractor func(r *http.Request) (string, error)

type Option struct {
	Issuer string
}
