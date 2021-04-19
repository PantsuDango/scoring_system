package db

import "scoring_system/model/tables"

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
