package db

import (
	"scoring_system/model/result"
	"scoring_system/model/tables"
)

type ScoringDB struct{}

func (ScoringDB) GetUserInfo(UserName string, Type int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("username = ? AND type = ? AND status= 0 ", UserName, Type).First(&user).Error
	return user, err
}

func (ScoringDB) QueryUserById(userId int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("id = ? AND status = 0", userId).First(&user).Error
	return user, err
}

func (ScoringDB) AddUser(user tables.User) error {
	err := exeDB.Create(&user).Error
	return err
}

func (ScoringDB) QueryAllUser() []result.ListUser {
	var user []result.ListUser
	exeDB.Where(`status = 0`).Find(&user)
	return user
}

func (ScoringDB) CreateProject(project *tables.Project) error {
	err := exeDB.Create(&project).Error
	return err
}

func (ScoringDB) CreateProjectUserMap(project_user_map *tables.ProjectUserMap) error {
	err := exeDB.Create(&project_user_map).Error
	return err
}

func (ScoringDB) SelectAllProject() []tables.Project {
	var project []tables.Project
	exeDB.Find(&project)
	return project
}

func (ScoringDB) SelectProjectUserMap(project_id int) []tables.ProjectUserMap {
	var project_user_map []tables.ProjectUserMap
	exeDB.Where(`project_id = ?`, project_id).Find(&project_user_map)
	return project_user_map
}

func (ScoringDB) SelectScore(project_id, player_id, judges_id int) tables.Score {
	var score tables.Score
	exeDB.Where(`project_id = ? AND player_id = ? AND judges_id = ?`, project_id, player_id, judges_id).Find(&score)
	return score
}

func (ScoringDB) SelectScore2(project_id, player_id int) ([]tables.Score, int) {
	var score []tables.Score
	var count int
	exeDB.Where(`project_id = ? AND player_id = ?`, project_id, player_id).Find(&score).Count(&count)
	return score, count
}

func (ScoringDB) SelectProjectUserMapByUserId(user_id int) []tables.ProjectUserMap {
	var project_user_map []tables.ProjectUserMap
	exeDB.Where(`user_id = ?`, user_id).Find(&project_user_map)
	return project_user_map
}

func (ScoringDB) SelectProjectByUserId(id int) tables.Project {
	var project tables.Project
	exeDB.Where(`id = ?`, id).Find(&project)
	return project
}

func (ScoringDB) SelectProjectUserMapToPlayer(project_id int) []tables.ProjectUserMap {
	var project_user_map []tables.ProjectUserMap
	exeDB.Where(`project_id = ? and type = 3`, project_id).Find(&project_user_map)
	return project_user_map
}

func (ScoringDB) CreateScore(score tables.Score) error {
	err := exeDB.Save(&score).Error
	return err
}

func (ScoringDB) ModifyUser(user tables.User) error {
	err := exeDB.Save(&user).Error
	return err
}
