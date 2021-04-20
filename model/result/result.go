package result

type ListUser struct {
	ID       int    `json:"ID"          form:"ID"          gorm:"column:id"`
	Nick     string `json:"Nick"        form:"Nick"        gorm:"column:nick"`
	Username string `json:"Username"    form:"Username"    gorm:"column:username"`
	Type     int    `json:"Type"        form:"Type"        gorm:"column:type"`
}

func (ListUser) TableName() string {
	return "user"
}
