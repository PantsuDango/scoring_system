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
	"sort"
	"strconv"
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
		ListProject.PlayerInfo = make([]result.PlayerInfo, 0)
		ListProject.JudgesInfo = make([]result.JudgesInfo, 0)

		ListProject.ID = project.ID
		ListProject.Content = project.Content
		ListProject.Name = project.Name
		ListProject.CreatedAt = project.CreatedAt.Format("2006-01-02 15:04:05")

		project_user_maps := Controller.ScoringDB.SelectProjectUserMap(project.ID)
		for _, project_user_map := range project_user_maps {
			if project_user_map.Type == 3 {
				var PlayerInfo result.PlayerInfo
				player, _ := Controller.ScoringDB.QueryUserById(project_user_map.UserId)
				PlayerInfo.ID = player.ID
				PlayerInfo.Username = player.Username
				PlayerInfo.Nick = player.Nick

				score, count := Controller.ScoringDB.SelectScore2(project.ID, player.ID)
				if count == 0 {
					PlayerInfo.Score = 0
				} else {
					var count int
					for _, val := range score {
						count += val.Score
					}
					fmt.Println(count)
					value := float64(count) / float64(len(score))
					value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
					PlayerInfo.Score = value
				}

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
		ListProject.PlayerInfo = SortByAge(ListProject.PlayerInfo)
		ListProjectResult = append(ListProjectResult, ListProject)
	} // 结束循环

	JSONSuccess(ctx, http.StatusOK, ListProjectResult)
}

func SortByAge(u []result.PlayerInfo) []result.PlayerInfo {

	sort.Slice(u, func(i, j int) bool {
		return u[i].Score > u[j].Score
	})

	return u
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

			score, count := Controller.ScoringDB.SelectScore2(project.ID, player.ID)
			if count == 0 {
				PlayerInfo.Score = 0
			} else {
				var count int
				for _, val := range score {
					count += val.Score
				}
				fmt.Println(count)
				value := float64(count) / float64(len(score))
				value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
				PlayerInfo.Score = value
			}

			Project.PlayerInfo = append(Project.PlayerInfo, PlayerInfo)
		}
		Project.PlayerInfo = SortByAge(Project.PlayerInfo)
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

	var err error
	for _, PlayerInfo := range ScoringParams.PlayerInfo {
		score := Controller.ScoringDB.SelectScore(ScoringParams.ProjectId, PlayerInfo.PlayerId, user.ID)
		score.ProjectId = ScoringParams.ProjectId
		score.PlayerId = PlayerInfo.PlayerId
		score.Score = PlayerInfo.Score
		score.JudgesId = user.ID
		err = Controller.ScoringDB.CreateScore(score)
	}
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

	var operator tables.OladUser
	if err := ctx.ShouldBindBodyWith(&operator, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid JSON or Illegal request parameter.", gin.H{
			"Code":    "InvalidJSON",
			"Message": err.Error(),
		})
		return
	}

	if len(operator.Password) > 0 {
		s := operator.OldPassword + user.Salt
		if fmt.Sprintf("%x", md5.Sum([]byte(s))) == user.Password {
			user.Salt = GetRandomString(8)
			s := operator.Password + user.Salt
			user.Password = fmt.Sprintf("%x", md5.Sum([]byte(s)))
		} else {
			JSONFail(ctx, http.StatusOK, PasswordError, "OldPassword error.", gin.H{
				"Code":    "InvalidJSON",
				"Message": "OldPassword error.",
			})
			return
		}
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
