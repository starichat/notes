package service

import "github.com/starichat/notes/DesignPatter/credit_sys/dao/credit"

type CreditService struct{}

// 赚取积分
/**
用户可以通过如下途径赚取积分：
1. 下订单，按照订单金额的10%获取10积分，比如100元的订单就可以获取10积分，不足1积分的不计入
2. 签到：签到一次获取5积分，每天只能签到一次
 */
func (c *CreditService) earnCredit(credit *credit.Credit){


}

// 消费积分
/**
用户在消费的时候可以通过积分抵扣
1. 消费的时候检查积分有效期，计算出还在有效期内的积分总额
2. 用户可以设置消费积分数量
3. 进行消费的时候，直接按照有效期从高到低进行消费
 */
func (c *CreditService) consumeCredit(credit *credit.Credit){

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