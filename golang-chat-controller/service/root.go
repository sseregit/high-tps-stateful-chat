package service

import (
	"golang-chat-controller/repository"
	"golang-chat-controller/types/table"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	s := &Service{repository: repository}
	return s
}

func (s *Service) GetAvailableServerList() ([]*table.Serverinfo, error) {
	return s.repository.GetAvailableServerList()
}
