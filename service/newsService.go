package service

import (
	"API/model/news"
	"API/repository"
	"time"
	newsViewModel "َAPI/ViewModel/news"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
	CreateNewUser(userInput newsViewModel.CreateNewsViewModel) (string, error)
}

type newsService struct {
}

func NewNewsService() NewsService {
	return newsService{}
}

func (newsService) GetNewsList() ([]news.News, error) {

	newsRepository := repository.NewNewsRepository()
	newsList, err := newsRepository.GetNewsList()
	return newsList, err
}
func (s newsService) CreateNewUser(userInput newsViewModel.CreateNewsViewModel) (string, error) {
	newsEntity := news.News{
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}

	newsRepository := repository.NewNewsRepository()
	newsId, err := newsRepository.InsertNews(newsEntity)

	return newsId, err
}
