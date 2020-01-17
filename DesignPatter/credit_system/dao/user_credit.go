package dao


// 新增用户积分,增加积分
func (d *Dao)  AddUserCredit(u *model.UserCredit) (err error) {
	return d.DB.Table("user_credit").Create(u).Error
}

// 更新用户积分
func (d *Dao) UpdateUserCredit(u *model.UserCredit) (err error) {
	return d.DB.Table("user_credit").Where("id = ?",u.ID).Update(u).Error
}

// 根据积分id删除用户积分，消费积分
func (d *Dao)  DeleteUserCredit(id int) (err error) {
	// 删除用户积分表的积分
	// 业务考虑，积分表也没有保留积分的必要，也应该一并删除
	err = d.DeleteCredit(id)
	if err != nil {
		log.Println("deletecredit error")
	}
	d.DB.Table("user_credit").Where("id = ?",u.ID).Delete(model.UserCredit{}).Error
	return err
	
}



