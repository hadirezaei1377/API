package utility

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// we must create a validator(customized)
type CustomValidator struct {
	Validator *validator.Validate
}

// we need to a function for doing validation
// this function recieves everything cause has interface
func (cv *CustomValidator) Validate(i interface{}) error {

	// recieve a struct and validate it
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}
