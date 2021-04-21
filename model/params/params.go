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
	ProjectId int `json:"ProjectId"  form:"ProjectId"  binding:"required"`
	PlayerId  int `json:"PlayerId"  form:"PlayerId"  binding:"required"`
	Score     int `json:"Score"  form:"Score"  binding:"required"`
}
