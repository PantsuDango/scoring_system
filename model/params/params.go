package params

type ModActIndex struct {
	Module string `json:"Module" form:"Module" binding:"required"`
	Action string `json:"Action" form:"Action" binding:"required"`
}

type AddProjectParams struct {
	Name     string `json:"Name"     form:"Name"     binding:"required"`
	Content  string `json:"Content"  form:"Content"`
	PlayerId []int  `json:"PlayerId" form:"PlayerId" binding:"required"`
	JudgesId []int  `json:"JudgesId" form:"JudgesId" binding:"required"`
}

type ScoringParams struct {
	ProjectId  int          `json:"ProjectId"  form:"ProjectId"  binding:"required"`
	PlayerInfo []PlayerInfo `json:"PlayerInfo"  form:"PlayerInfo"  binding:"required"`
}

type PlayerInfo struct {
	PlayerId int `json:"ID"  form:"ID"`
	Score    int `json:"Score"  form:"Score"`
}
