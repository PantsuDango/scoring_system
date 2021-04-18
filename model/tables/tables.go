package tables

import "time"


type User struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	Nick      string    `json:"Nick"        form:"Nick"        gorm:"column:nick"`
	Username  string    `json:"Username"    form:"Username"    gorm:"column:username"`
	Password  string    `json:"Password"    form:"Password"    gorm:"column:password"`
	Salt      string    `json:"Salt"        form:"Salt"        gorm:"column:salt"`
	Type      int       `json:"Type"        form:"Type"        gorm:"column:type"`
	Status    int       `json:"Status"      form:"Status"      gorm:"column:status"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}