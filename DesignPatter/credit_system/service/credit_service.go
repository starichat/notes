package service

type CreditService struct{}

type req struct {
	UserId uint64
	ChannelId string
	EventId string
}

//// 赚取积分
///**
//用户可以通过如下途径赚取积分：
//1. 下订单，按照订单金额的10%获取10积分，比如100元的订单就可以获取10积分，不足1积分的不计入
//2. 签到：签到一次获取5积分，每天只能签到一次
// */
//func (service *Service) EarnCredit(req *req)  (map[string]interface{}) {
//	credit := &model.Credit{}
//	if credit.ChannelId.GetType() == 0 { // 根据订单金额
//		credit.credit = credit. EventId.GetMoney()/ 10
//	} else { // 签到
//		credit.credit = 5
//	}
//	credit = req
//	err ,_ := s.dao.Add()
//	return err
//}
//
//// 消费积分
///**
//用户在消费的时候可以通过积分抵扣或者兑换优惠卷，这里暂时只考虑一个消费劵
//1. 消费的时候检查积分有效期，计算出还在有效期内的积分总额
//2. 用户可以设置消费积分数量
//3. 进行消费的时候，直接按照有效期从高到低进行抵扣
// */
//func (service *Service) consumeCredit(req,sum) credit.Credit {
//	sum = credit.credit
//
//}
// 获取积分
/**
获取用户可用总积分
 */
func GetCreditPoints(){}

/**
获取用户积分有效期等细节
 */
func getCreditDetail(){}

/**
获取用户获取的积分
 */
func getDetailByEarn(){}

/**
获取用户消费的积分
 */
func getDetailByConsume(){}