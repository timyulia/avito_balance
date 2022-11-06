package service

import (
	"balance"
	"balance/pkg/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) MakeReport(year, month int) error {

	return s.repo.MakeReport(year, month)
}

func (s *InfoService) GiveName(serv balance.Report) error {

	return s.repo.GiveName(serv)
}
