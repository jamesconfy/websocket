package service

import "project-name/internal/se"

type HomeService interface {
	CreateHome() (string, *se.ServiceError)
}

type homeSrv struct{}

func (h *homeSrv) CreateHome() (string, *se.ServiceError) {
	return "Home", nil
}

func NewHomeService() HomeService {
	return &homeSrv{}
}
