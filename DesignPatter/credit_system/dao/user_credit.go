package dao

import (
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"log"
)

// 新增用户积分,增加积分
func (d *Dao)  AddUserCredit(u *model.UserCredit) (err error) {
	return d.DB.Table("user_credit").Create(u).Error
}

// 更新用户积分
func (d *Dao) UpdateUserCredit(u *model.UserCredit) (err error) {
	return d.DB.Table("user_credit").Where("id = ?",u.Id).Update(u).Error
}

// 查找用户积分信息
func (d *Dao) FindCreditsById(userID,limit,offset int) (creditIds []int,err error) {
	err= d.DB.Table("user_credit").Select("credit_id").Where("user_id = ?",userID).Limit(limit).Offset(offset).Find(&creditIds).Error
	return
}

// 根据积分id删除用户积分，消费积分
func (d *Dao)  DeleteUserCredit(id uint) (err error) {
	// 删除用户积分表的积分
	// 业务考虑，积分表也没有保留积分的必要，也应该一并删除
	err = d.DeleteCredit(id)
	if err != nil {
		panic("deletecredit error:"+err.Error())
		log.Println("deletecredit error")
	}
	err = d.DB.Table("user_credit").Where("credit_id = ?",id).Delete(model.UserCredit{}).Error
	return err
	
}



