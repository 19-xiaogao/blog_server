package service

import (
	"xiaolong_blog/internal/model"
	"xiaolong_blog/pkg/app"
)

type CreateArticleRequest struct {
	Title    string `form:"title" binding:"required"`
	Describe string `form:"describe" binding:"required"`
	Content  string `form:"content" binding:"required"`
	MusicUrl string `form:"musicUrl" binding:"required"`
	CoverUrl string `form:"coverUrl" binding:"required"`
}

type DeleteArticleRequest struct {
	ID int `form:"id" binding:"required" uri:"id"`
}

type QueryArticleRequest struct {
	DeleteArticleRequest
}

type UpdateArticleRequest struct {
	CreateArticleRequest
	ID int `form:"id" binding:"required"`
}
type ListArticleRequest struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

func (svc *Service) CreateArticle(params *CreateArticleRequest) error {
	return svc.dao.CreateArticle(params.Title, params.Describe, params.Content, params.CoverUrl, params.MusicUrl)
}

func (svc *Service) DeleteArticle(params *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(params.ID)
}

func (svc *Service) QueryArticle(params *QueryArticleRequest) interface{} {
	return svc.dao.QueryArticle(params.ID)
}
func (svc *Service) UpdateArticle(params *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(params.ID, params.Title, params.Describe, params.Content, params.CoverUrl, params.MusicUrl)
}
func (svc *Service) ListArticle(params *app.Pager) ([]*model.Article, error) {
	return svc.dao.ListArticle(params.Page, params.PageSize)
}

func (svc *Service) CountArticle() (int, error) {
	return svc.dao.CountArticle()
}
