package service

import "github.com/ShekleinAleksey/gomessage/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
