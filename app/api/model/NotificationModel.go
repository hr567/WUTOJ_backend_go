package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"time"
)

type Notification struct {
	ID         int       `json:"id" form:"id"`
	Title      string    `json:"title" form:"title"`
	Content    string    `json:"content" form:"content"`
	SubmitTime time.Time `json:"submit_time" form:"submit_time"`
	ModifyTime time.Time `json:"modify_time" form:"modify_time"`
	ContestID  int       `json:"contest_id" form:"contest_id"`
	UserID     int       `json:"user_id" form:"user_id"`
	Status     int       `json:"status" form:"status"`
	EndTime    int       `json:"end_time" form:"end_time"`
}

func (model *Notification) GetAllNotification() helper.ReturnType {
	var notifications []Notification

	err := db.Find(&notifications).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取成功", Data: notifications}
	}

}

func (model *Notification) GetNotificationByID(id int) helper.ReturnType {
	var notification Notification

	err := db.Where("id = ?", id).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取成功", Data: notification}
	}

}

func (model *Notification) GetNotification(ContestID int) helper.ReturnType {
	var notification = Notification{}

	err := db.Where("contest_id = ?", ContestID).Where("status = ?", 0).First(&notification).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "获取失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "获取成功", Data: notification}
	}

}
