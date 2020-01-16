package model

import "time"

// CreditInfo 积分信息
type CreditInfo struct {
	Id uint64 `"gorm:column:id;type:int;primary_key;AUTO_INCREMENT"`
	ChannelId string `"gorm:column:channel_id;type:varchar(32);"`
	EventId string  `"gorm:column:event_id;type:varchar(32);"`
	Credit int64  `"gorm:column:credit;type:varchar(32);"`
	CreatedTime time.Time  `"gorm:column:created_time;type:datetime;"`
	ExpiredTime time.Time  `"gorm:column:expired_time;type:datetime;"`
}

// UserInfo 用户信息
type UserInfo struct {
	Id uint64  `"gorm:column:id;type:int;primary_key;AUTO_INCREMENT"`
	Password string  `"gorm:column:password;type:varchar(32);"`
}

// UserCredit 用户积分信息
type UserCredit struct {
	Id uint64 `"gorm:column:id;type:int;primary_key;AUTO_INCREMENT"`
	UserId uint64 `"gorm:column:userid;type:int"`
	CreditId uint64 `"gorm:column:creditid;type:int"`
}