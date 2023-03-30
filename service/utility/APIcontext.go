package utility

/*

// for custom context ... c.any = usecase of context
import (
	"API/ViewModel/common/security"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// do a work that when I want ,I have this ability to add things to my context... like add or remove methods from context
type ApiContext struct {
	echo.Context
}

// add this method
func (c ApiContext) GetUserId() (userId string, err error) {
	defer func() {
		if r := recover(); r != nil {
			userId = ""
			err = errors.New("user is not login")
		}
	}()
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*security.JwtClaims)
	return claim.UserId, nil
}

// do this commands and this additional methods as a middleware in main.go

*/
