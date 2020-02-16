package service

import (
	_ "github.com/starichat/notes/DesignPatter/credit_system/dao"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	_ "github.com/starichat/notes/DesignPatter/credit_system/pkg"
	"time"

	_ "time"
)





// 赚取积分
/**
用户可以通过如下途径赚取积分：
1. 下订单，按照订单金额的10%获取10积分，比如100元的订单就可以获取10积分，不足1积分的不计入
2. 签到：签到一次获取5积分，每天只能签到一次
*/
func (s *Service) EarnCredit(reqCredit *model.CreditReq) (creditId int) {
	credit := &model.CreditInfo{
		ChannelId:   reqCredit.ChannelId,
		EventId:     reqCredit.EventId,
		Credit:      reqCredit.Credit,
		CreatedTime: time.Now(),
		ExpiredTime: reqCredit.ExpiredTime,
	}
	err := s.dao.AddCredit(credit)
	if err != nil {
		return -1
	}


	return 0;

}



// 消费积分
/**
用户在消费的时候可以通过积分抵扣或者兑换优惠卷，这里暂时只考虑一个消费劵
1. 消费的时候检查积分有效期，计算出还在有效期内的积分总额
2. 用户可以设置消费积分数量
3. 进行消费的时候，直接按照有效期从高到低进行抵扣
*/
func (s *Service) consumeCredit(reqCredit *model.CreditReq) (creditId int){
	return

}
// 根据列出所有积分
func (s *Service) GetCreditByCreditID(creditId int) (credit *model.CreditInfo) {
	credit, err := s.dao.FindCreditById(creditId)
	if err != nil {
	}
	return credit

}

/**
获取用户有效积分详情
 */
func (s *Service) getCreditDetail(userid, limit, offset int) (credits []*model.CreditInfo,sum int ){
	sum = 0
	ids, err := s.dao.FindCreditsById(userid,limit,offset)
	if err != nil {
		panic(err)
	}
	// for 循环获取总积分
	for _, id := range ids {
		credit := s.GetCreditByCreditID(id)
		if credit.Credit > 0{
			sum += credit.Credit
			credits = append(credits, credit)
		}
	}
	return credits,sum
}

func (s *Service) getCreditByUserId(userID int) (sum int) {
	return
}


func (s *Service) getTotalCredits(userid, limit, offset int)  (credits []*model.CreditInfo,sum int ) {
	return
}

func (s *Service) getEarnCredits(userid, limit, offset int) (credits []*model.CreditInfo,sum int ) {
	return
}

func (s *Service) getConsumeCredits(userid, limit, offset int) (credits []*model.CreditInfo,sum int ) {
	return
}



