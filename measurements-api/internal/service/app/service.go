package app

import (
	def "measurements-api/internal/service"
)

var _ def.AppService = (*service)(nil)

type service struct{}

func NewService() *service {
	return &service{}
}

func (u *service) IsOnline() bool {
	return true
}
