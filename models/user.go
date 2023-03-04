package models

import (
	"github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
)

type Users struct {
	gorm.Model
	Mid        string         `json:"mid" gorm:"mid"`
	Password   string         `json:"password" gorm:"password"`
	Name       string         `json:"name" gorm:"name"`
	Avatar     string         `json:"avatar" gorm:"avatar"`
	TotalLogin int64          `json:"total" gorm:"total_login"`
	Equipments pq.StringArray `gorm:"type:varchar(255)[]"`
}

func GetUserByMid(mid string) (*Users, error) {
	var user Users
	result := db.Where("mid = $1", mid).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if result.Error == gorm.ErrRecordNotFound && result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func CreateUser(user *Users) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *Users) error {
	return db.Model(&Users{}).Where("mid = ?", user.Mid).Updates(user).Error
}

func UpdateUserByMap(data map[string]interface{}) error {
	if err := db.Model(&Users{}).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddUserLoginTimes(mid string, times int) error {
	return db.Model(&Users{}).Where("mid = ?", mid).Update("total_login", gorm.Expr("total_login + ?", times)).Error

}

func GetAllUser() ([]Users, error) {
	var users []Users
	result := db.Find(&users)
	return users, result.Error
}
