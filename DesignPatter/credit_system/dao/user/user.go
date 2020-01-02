package user

import "github.com/starichat/notes/DesignPatter/credit_sys/dao"

type User struct {
	ID uint `gorm:"column:id;type:int;primary_key;auto_increment"`
	UserName string `gorm:"column:username;type:varchar(32)"`
	CreditId uint64 `gorm:"column:credit_id;type:int;"`
}

// 用户

func AddUserCredit(u *User) {
	dao.DB.Create(u)
}

// 更新用户积分
func UpdateUserCredit(u *User) {
	dao.DB.Where("id = ?",u.ID).Update(u)
}

// 删除用户关联
func DeleteUserCredit() {

}

