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

type Project struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	Name      string    `json:"Name"        form:"Name"        gorm:"column:name"`
	Content   string    `json:"Content"     form:"Content"     gorm:"column:content"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (Project) TableName() string {
	return "project"
}

type ProjectUserMap struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	ProjectId int       `json:"ProjectId"   form:"ProjectId"   gorm:"column:project_id"`
	UserId    int       `json:"UserId"      form:"UserId"      gorm:"column:user_id"`
	Type      int       `json:"Type"        form:"Type"        gorm:"column:type"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (ProjectUserMap) TableName() string {
	return "project_user_map"
}

type Score struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	ProjectId int       `json:"ProjectId"   form:"ProjectId"   gorm:"column:project_id"`
	PlayerId  int       `json:"UserId"      form:"UserId"      gorm:"column:player_id"`
	Score     int       `json:"Score"       form:"Score"       gorm:"column:score"`
	JudgesId  int       `json:"JudgesId"    form:"JudgesId"    gorm:"column:judges_id"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (Score) TableName() string {
	return "score"
}
