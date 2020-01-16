package service


//// Service struct
//type Service struct {
//	c         *config.Config
//	dao       *dao.Dao
//}
//
//// New init
//func New(c *conf.Config) (s *Service) {
//	s = &Service{
//		c:         c,
//		dao:       dao.New(c),
//	}
//	return s
//}
//
//// Ping check db live
//func (s *Service) Ping(c context.Context) (err error) {
//	return s.dao.Ping(c)
//}