package security

// for work with JWT

import "github.com/golang-jwt/jwt"

// we want to have username in JWT, define it in this struct
// a series of claims are related to own JWTs and their management(standard claims)
type JwtClaims struct {
	UserName           string `json:"username"`
	jwt.StandardClaims        // include : subject , ID ...as defualt exist in jwt claims
}

// initialize JWTclaims in usercontroller
