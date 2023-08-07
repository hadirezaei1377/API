package controller

import (
	"API/Utility"
	"API/ViewModel/common/httpResponse"
	"API/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsController interface {
	GetNewsList(c echo.Context) error
}

type newsController struct {
}

func NewNewsController() NewsController {
	return newsController{}
}

func (nc newsController) GetNewsList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	newsService := service.NewNewsService()
	newsList, err := newsService.GetNewsList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}
