package service

import (
	_ "github.com/starichat/notes/DesignPatter/credit_system/dao"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"github.com/starichat/notes/DesignPatter/credit_system/pkg"
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
func (s *Service) EarnCredit(reqCredit *model.CreditReq) (int) {
	credit := &model.CreditInfo{
		Id: pkg.GenerateID(),
		ChannelId:   reqCredit.ChannelId,
		EventId:     reqCredit.EventId,
		Credit:      reqCredit.Credit,
		CreatedTime: time.Now(),
		ExpiredTime: reqCredit.ExpiredTime,
	}
	userCredit := &model.UserCredit{
		UserId:   reqCredit.UserId,
		CreditId: credit.Id,
	}
	err := s.dao.AddCredit(credit)
	err = s.dao.AddUserCredit(userCredit)
	if err != nil {
		return -1
	}
	return userCredit.CreditId;

}



// 消费积分
/**
用户在消费的时候可以通过积分抵扣或者兑换优惠卷，这里暂时只考虑一个消费劵
1. 消费的时候检查积分有效期，计算出还在有效期内的积分总额
2. 用户可以设置消费积分数量
3. 进行消费的时候，直接按照有效期从高到低进行抵扣
*/
func (s *Service) ConsumeCredit(reqCredit *model.CreditReq) (creditId int){
	// 1.计算所有积分
	// 2. 根据消费金额以及用户输入的积分数量，进行抵扣
	// 3. 根据有效期排序，输出积分对应的id，并从userCredit中删除
	// 4。 新增一条积分，消费完毕


	return

}
// 根据列出所有积分
func (s *Service) GetCreditByCreditID(userID int) (credit *model.CreditInfo) {
	credit, err := s.dao.FindCreditById(userID)
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



func (s *Service) getTotalCredits(userid, limit, offset int)  (sum int ) {
	sum = 0
	ids, err := s.dao.FindCreditsById(userid,limit,offset)
	if err != nil {
		panic(err)
	}
	// for 循环获取总积分 ,筛选过期积分
	for _, id := range ids {
		credit := s.GetCreditByCreditID(id)
		sum+=credit.Credit
	}
	return sum
}

func (s *Service) getEarnCredits(userid, limit, offset int) (credits []*model.CreditInfo) {

	ids, err := s.dao.FindCreditsById(userid,limit,offset)
	if err != nil {
		panic(err)
	}
	// for 循环获取总积分 ,筛选过期积分
	for _, id := range ids {
		if pkg.GetType(string(id))==1 {
			credit := s.GetCreditByCreditID(id)
			credits = append(credits, credit)
		}

	}
	return credits
}

func (s *Service) getConsumeCredits(userid, limit, offset int) (credits []*model.CreditInfo) {

	ids, err := s.dao.FindCreditsById(userid,limit,offset)
	if err != nil {
		panic(err)
	}
	// for 循环获取总积分 ,筛选过期积分
	for _, id := range ids {
		if pkg.GetType(string(id))==0 {
			credit := s.GetCreditByCreditID(id)
			credits = append(credits, credit)
		}

	}
	return credits
}



