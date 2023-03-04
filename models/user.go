package models

import (
	"pp-bakcend/pkg/logging"

	"github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Mid        string         `json:"mid" gorm:"mid"`
	Password   string         `json:"password" gorm:"password"`
	Name       string         `json:"name" gorm:"name"`
	Avatar     string         `json:"avatar" gorm:"avatar"`
	TotalLogin int64          `json:"total" gorm:"total_login"`
	Equipments pq.StringArray `gorm:"type:varchar(255)[]"`
}

func GetUserByMid(mid string) (*User, error) {
	var user User
	result := db.Where("mid = $1", mid).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if result.Error == gorm.ErrRecordNotFound && result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func CreateUser(user *User) error {
	logging.Trace(user)
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) error {
	return db.Model(&User{}).Where("mid = ?", user.Mid).Updates(user).Error
}

func UpdateUserByMap(data map[string]interface{}) error {
	if err := db.Model(&User{}).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddUserLoginTimes(mid string, times int) error {
	return db.Model(&User{}).Where("mid = ?", mid).Update("total_login", gorm.Expr("total_login + ?", times)).Error

}

func GetAllUser() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}
