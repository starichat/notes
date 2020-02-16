package dao

import (
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"log"
	_ "time"
)


// 插入数据
func (d *Dao) AddCredit(c *model.CreditInfo) (err error) {
	return d.DB.Table("credits").Create(c).Error
}

// 更新数据
func (d *Dao) UpdateCredit(c *model.CreditInfo) (err error) {
	log.Println(c)
	return d.DB.Table("credits").Where("id = ?",c.Id).Update(c).Error
}

// 查询指定id的详细内容
func (d *Dao) FindCreditById(id int) (credit *model.CreditInfo, err error) {
	credit = &model.CreditInfo{}
	err = d.DB.Table("credits").Select("id, channel_id, event_id, credit, created_time, expired_time").Where("id = ?",id).Find(credit).Error
	if err != nil {
		log.Printf("查询数据失败： %v",err)
		return nil, err
	}
	return credit, err
}

// 分页查询
func (d *Dao) FindLimitCredit(limit, offset uint)(credits []*model.CreditInfo, err error) {
	err = d.DB.Table("credits").Select("id, channel_id, event_id, credit, created_time, expired_time").Limit(limit).Offset(offset).Find(&credits).Error
	if err != nil {
		log.Printf("查询数据失败： %v",err)
		return nil,err
	}
	return credits, err
}

// 删除数据
func (d *Dao) DeleteCredit(id uint) (err error) {
	err = d.DB.Table("credits").Where("id = ?",id).Delete(model.CreditInfo{}).Error
	if err != nil {
		log.Printf("删除失败:   %v",err)
		return err
	}
	return nil
}