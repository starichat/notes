package service

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/dao"
	"github.com/starichat/notes/DesignPatter/credit_sys/dao/credit"

	"time"
)

type CreditService struct{}

type UserCredit struct {
	UserId uint64
	ChannelId string
	EventId string
	Credit int64
	CreatedTime time.Time
	ExpiredTime time.Time
}

// 赚取积分
/**
用户可以通过如下途径赚取积分：
1. 下订单，按照订单金额的10%获取10积分，比如100元的订单就可以获取10积分，不足1积分的不计入
2. 签到：签到一次获取5积分，每天只能签到一次
 */
func (c *CreditService) EarnCredit(userCredit UserCredit) credit.Credit {
	//if userCredit.ChannelId == 0 {
	//	userCredit.Credit = 100 / 10
	//} else {
	//	userCredit.Credit = 5
	//}

	return credit.Credit{
		Id:          1424,
		ChannelId:   userCredit.ChannelId,
		EventId:     userCredit.EventId,
		Credit:      userCredit.Credit,
		CreatedTime: userCredit.CreatedTime,
		ExpiredTime: userCredit.ExpiredTime,
	}

}

// 消费积分
/**
用户在消费的时候可以通过积分抵扣或者兑换优惠卷，这里暂时只考虑一个消费劵
1. 消费的时候检查积分有效期，计算出还在有效期内的积分总额
2. 用户可以设置消费积分数量
3. 进行消费的时候，直接按照有效期从高到低进行抵扣
 */
func (c *CreditService) consumeCredit(userCredit UserCredit) credit.Credit {
	return credit.Credit{
		Id:          1424,
		ChannelId:   userCredit.ChannelId,
		EventId:     userCredit.EventId,
		Credit:      userCredit.Credit,
		CreatedTime: userCredit.CreatedTime,
		ExpiredTime: userCredit.ExpiredTime,
	}
}
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