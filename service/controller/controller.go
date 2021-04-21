package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"scoring_system/model/params"
	"scoring_system/model/result"
	"scoring_system/model/tables"
)

// 查询所有用户
func (Controller Controller) ListUser(ctx *gin.Context, user tables.User) {
	// 如果不是主账号
	if user.Type != 1 {
		JSONFail(ctx, http.StatusOK, AccessDeny, "user type error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "user type error",
		})
		return
	}

	user_info := Controller.ScoringDB.QueryAllUser()

	JSONSuccess(ctx, http.StatusOK, user_info)
}

// 新建项目
func (Controller Controller) AddProject(ctx *gin.Context, user tables.User) {

	var AddProjectParams params.AddProjectParams
	if err := ctx.ShouldBindBodyWith(&AddProjectParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	// 如果不是主账号
	if user.Type != 1 {
		JSONFail(ctx, http.StatusOK, AccessDeny, "user type error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "user type error",
		})
		return
	}
	// 校验前端传来的参数是否符合预期
	if len(AddProjectParams.Name) == 0 || len(AddProjectParams.PlayerId) == 0 || len(AddProjectParams.JudgesId) == 0 {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "Invalid JSON or Illegal request parameter.",
		})
		return
	}

	// 创建项目
	var project tables.Project
	project.Name = AddProjectParams.Name
	project.Content = AddProjectParams.Content
	err := Controller.ScoringDB.CreateProject(&project)

	if err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	// 绑定项目和选手的关系
	for _, PlayerId := range AddProjectParams.PlayerId {
		var project_user_map tables.ProjectUserMap
		project_user_map.ProjectId = project.ID
		project_user_map.UserId = PlayerId
		project_user_map.Type = 3
		user, _ := Controller.ScoringDB.QueryUserById(PlayerId)
		if user.Type != 3 {
			continue
		}
		err = Controller.ScoringDB.CreateProjectUserMap(&project_user_map)
	}

	// 绑定项目和评委的关系
	for _, JudgesId := range AddProjectParams.JudgesId {
		var project_user_map tables.ProjectUserMap
		project_user_map.ProjectId = project.ID
		project_user_map.UserId = JudgesId
		project_user_map.Type = 2
		user, _ := Controller.ScoringDB.QueryUserById(JudgesId)
		if user.Type != 2 {
			continue
		}
		err = Controller.ScoringDB.CreateProjectUserMap(&project_user_map)
	}

	if err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 查询所有项目
func (Controller Controller) ListProject(ctx *gin.Context, user tables.User) {

	// 如果不是主账号
	if user.Type != 1 {
		JSONFail(ctx, http.StatusOK, AccessDeny, "user type error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "user type error",
		})
		return
	}

	var ListProjectResult []result.ListProject
	ListProjectResult = make([]result.ListProject, 0)

	projects := Controller.ScoringDB.SelectAllProject()
	for _, project := range projects {
		var ListProject result.ListProject

		ListProject.ID = project.ID
		ListProject.Content = project.Content
		ListProject.Name = project.Name
		ListProject.CreatedAt = project.CreatedAt.Format("2006-01-02 15:04:05")

		project_user_maps := Controller.ScoringDB.SelectProjectUserMap(project.ID)
		for _, project_user_map := range project_user_maps {
			if project_user_map.Type == 2 {
				var PlayerInfo result.PlayerInfo
				player, _ := Controller.ScoringDB.QueryUserById(project_user_map.UserId)
				PlayerInfo.ID = player.ID
				PlayerInfo.Username = player.Username
				PlayerInfo.Nick = player.Nick
				score := Controller.ScoringDB.SelectScore(project.ID, player.ID)
				PlayerInfo.Score = score.Score
				ListProject.PlayerInfo = append(ListProject.PlayerInfo, PlayerInfo)
			} else {
				var JudgesInfo result.JudgesInfo
				judges, _ := Controller.ScoringDB.QueryUserById(project_user_map.UserId)
				JudgesInfo.ID = judges.ID
				JudgesInfo.Username = judges.Username
				JudgesInfo.Nick = judges.Nick
				ListProject.JudgesInfo = append(ListProject.JudgesInfo, JudgesInfo)
			}
		} // 循环结束

		ListProjectResult = append(ListProjectResult, ListProject)
	} // 结束循环

	JSONSuccess(ctx, http.StatusOK, ListProjectResult)
}

// 新建项目
func (Controller Controller) ProjectInfo(ctx *gin.Context, user tables.User) {

	var ProjectInfoResult []result.ProjectInfo

	project_user_maps := Controller.ScoringDB.SelectProjectUserMapByUserId(user.ID)
	for _, project_user_map := range project_user_maps {
		// 查询项目信息
		var Project result.ProjectInfo
		project := Controller.ScoringDB.SelectProjectByUserId(project_user_map.ProjectId)
		Project.ID = project.ID
		Project.Name = project.Name
		Project.Content = project.Content
		Project.CreatedAt = project.CreatedAt.Format("2006-01-02 15:04:05")
		// 查询选手信息
		tmps := Controller.ScoringDB.SelectProjectUserMapToPlayer(project.ID)
		for _, tmp := range tmps {
			var PlayerInfo result.PlayerInfo
			player, _ := Controller.ScoringDB.QueryUserById(tmp.UserId)
			PlayerInfo.ID = player.ID
			PlayerInfo.Username = player.Username
			PlayerInfo.Nick = player.Nick
			score := Controller.ScoringDB.SelectScore(project.ID, player.ID)
			PlayerInfo.Score = score.Score
			Project.PlayerInfo = append(Project.PlayerInfo, PlayerInfo)
		}
		ProjectInfoResult = append(ProjectInfoResult, Project)
	}

	JSONSuccess(ctx, http.StatusOK, ProjectInfoResult)
}

// 打分
func (Controller Controller) Scoring(ctx *gin.Context, user tables.User) {

	var ScoringParams params.ScoringParams
	if err := ctx.ShouldBindBodyWith(&ScoringParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	// 如果不是主账号
	if user.Type != 2 {
		JSONFail(ctx, http.StatusOK, AccessDeny, "user type error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "user type error",
		})
		return
	}

	var score tables.Score
	score.ProjectId = ScoringParams.ProjectId
	score.PlayerId = ScoringParams.ProjectId
	score.Score = ScoringParams.Score
	score.JudgesId = user.ID
	err := Controller.ScoringDB.CreateScore(score)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDeny, "create score fail.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 打分
func (Controller Controller) ModifyUser(ctx *gin.Context, user tables.User) {

	var operator tables.User
	if err := ctx.ShouldBindBodyWith(&operator, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	if len(operator.Password) > 0 {
		user.Salt = GetRandomString(8)
		s := operator.Password + user.Salt
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(s)))
	}

	if len(operator.Nick) > 0 {
		user.Nick = operator.Nick
	}

	err := Controller.ScoringDB.ModifyUser(user)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDeny, "update user fail.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}
