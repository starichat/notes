package model

type Attendance struct {
	UserId     int `form:"userId" json:"userId" binding:"required"`
	EventType string `form:"eventType" json:"eventType" binding:"required"`
}

type Consume struct {
	UserId	string `form:"userId" json:"userId" binding:"required"`
	EventType string `form:"eventType" json:"eventType" binding:"required"`
	Amount	int `form:"amount" json:"amount" binding:"required"`
}
