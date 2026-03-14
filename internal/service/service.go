package service

import (
	"context"
	"encoding/json"
	"microservice-go/internal/model"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*model.CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{url: url}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*model.CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fact := &model.CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
