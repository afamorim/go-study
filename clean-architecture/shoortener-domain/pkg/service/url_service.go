package service

import (
	"errors"

	"../model"
)

type UrlRepository interface {
	Save(url model.Url) (model.Url, error)
}

type UrlService struct {
	urlRepository UrlRepository
}

func (s *UrlService) save(url *model.Url) (interface{}, error) {
	if url.OriginalUrl == "" {
		return nil, errors.New("Original URL is mandatory")
	}
	return nil, nil
}
