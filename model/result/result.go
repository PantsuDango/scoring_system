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

type ListProject struct {
	ID         int          `json:"ID"          form:"ID"`
	Name       string       `json:"Name"        form:"Name"`
	Content    string       `json:"Content"     form:"Content"`
	CreatedAt  string       `json:"CreateTime"  form:"CreateTime"`
	PlayerInfo []PlayerInfo `json:"PlayerInfo"  form:"PlayerInfo"`
	JudgesInfo []JudgesInfo `json:"JudgesInfo"  form:"JudgesInfo"`
}

type PlayerInfo struct {
	ID       int     `json:"ID"        form:"ID"`
	Nick     string  `json:"Nick"      form:"Nick"`
	Username string  `json:"Username"  form:"Username"`
	Score    float64 `json:"Score"     form:"Score"`
}

type JudgesInfo struct {
	ID       int    `json:"ID"       form:"ID"`
	Nick     string `json:"Nick"     form:"Nick"`
	Username string `json:"Username" form:"Username"`
}

type ProjectInfo struct {
	ID         int          `json:"ID"          form:"ID"`
	Name       string       `json:"Name"        form:"Name"`
	Content    string       `json:"Content"     form:"Content"`
	CreatedAt  string       `json:"CreateTime"  form:"CreateTime"`
	PlayerInfo []PlayerInfo `json:"PlayerInfo"  form:"PlayerInfo"`
}
