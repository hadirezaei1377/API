package controller

import (
	"API/Utility"
	"API/ViewModel/common/httpResponse"
	"API/service"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
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
func (nc newsController) CreateNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	newNews := new(newsViewModel.CreateNewsViewModel)

	if err := apiContext.Bind(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		fileName := uuid.New().String() + filepath.Ext(file.Filename)

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		newNews.ImageName = fileName
	} else {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("image not found"))
	}

	newsService := service.NewNewsService()
	newNewsId, err := newsService.CreateNewUser(*newNews)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newNewsId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}
