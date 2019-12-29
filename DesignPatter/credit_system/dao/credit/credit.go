package credit

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/dao"
	"log"
)

type Credit struct {
	id uint
	channelId string
	eventId string
	credit string
	createTime string
	expiredTime string
}

// 插入数据
func  addCredit() (err error) {
	c := &Credit{
		id:          1223,
		channelId:   "133",
		eventId:     "1424",
		credit:      "1412",
		createTime:  "3331323",
		expiredTime: "1313",
	}
	log.Println(&c)
	if err = dao.DB.Create(c).Error; err != nil {
		log.Printf("微博创建失败: %v", err)
		return err
	}

	log.Println("add")
	return nil
}

// 更新数据
func updateCredit(c *Credit) error{
	return dao.DB.Where("id = ?",c.id).Update(c).Error
}

// 查询数据
func findTotalCredit(id uint) (credits []Credit, err error) {
	err = dao.DB.Select("id, channelId, eventId, credit, createTime, expiredTime").Find(&credits).Error
	return credits,err
}

// 分页查询
func findLimitCredit(id uint) Credit {
	return Credit{
		id:          0,
		channelId:   "",
		eventId:     "",
		credit:      "",
		createTime:  "",
		expiredTime: "",
	}
}