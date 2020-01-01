package credit

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/dao"
	"log"
	"time"
)

type Credit struct {
	Id uint64 `"gorm:column:id;type:int;primary_key;AUTO_INCREMENT"`
	ChannelId string `"gorm:column:channel_id;type:varchar(32);"`
	EventId string  `"gorm:column:event_id;type:varchar(32);"`
	Credit string  `"gorm:column:credit;type:varchar(32);"`
	CreatedTime time.Time  `"gorm:column:created_time;type:datetime;"`
	ExpiredTime time.Time  `"gorm:column:expired_time;type:datetime;"`
}

// 插入数据
func (c *Credit) AddCredit() (err error) {

	if err = dao.DB.Create(&c).Error; err != nil {
		log.Printf("数据创建失败: %v", err)
		return err
	}
	return nil
}

// 更新数据
func UpdateCredit(c *Credit) (err error) {
	log.Println(c.Id)
	if err = dao.DB.Table("credits").Where("id = ?",c.Id).Update(c).Error; err != nil {
		log.Printf("数据更新失败: %v", err)
		return err
	}
	return nil
}

// 查询指定id数据
func FindCreditById(id uint) (credit *Credit, err error) {
	err = dao.DB.Select("id, channel_id, event_id, credit, created_time, expired_time").Where("id = ?",id).Find(&credit).Error
	if err != nil {
		log.Printf("查询数据失败： %v",err)
		return nil,err
	}
	return credit, err
}

// 分页查询
func FindLimitCredit(limit, offset uint)(credits []Credit, err error) {
	err = dao.DB.Select("id, channel_id, event_id, credit, created_time, expired_time").Limit(limit).Offset(offset).Find(&credits).Error
	if err != nil {
		log.Printf("查询数据失败： %v",err)
		return nil,err
	}
	return credits, err
}

// 删除数据
func DeleteCredit(id uint) (err error) {
	err = dao.DB.Where("id = ?",id).Delete(Credit{}).Error
	if err != nil {
		log.Printf("删除失败:   %v",err)
		return err
	}
	return nil
}