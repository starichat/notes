package service

import (
	"context"
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"github.com/starichat/notes/DesignPatter/credit_system/dao"
)

// Service struct
type Service struct {
	c         *config.DBConfig
	dao       *dao.Dao
}

// New init
func New(c *config.DBConfig) (s *Service) {
	s = &Service{
		c:         c,
		dao:       dao.New(c),
	}
	return s
}

// Ping check db live
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}