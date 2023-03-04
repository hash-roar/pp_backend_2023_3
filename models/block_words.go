package models

import "github.com/jinzhu/gorm"

type BlockWords struct {
	gorm.Model
	Mid     string
	Word    string
	Handle  string
	Visible bool
}

func AddBlockWords(mid string, word string, handle string) error {

	return db.Create(&BlockWords{
		Mid:    mid,
		Word:   word,
		Handle: handle,
	}).Error
}

func DeleteBlockWord(mid string, word string) error {
	return db.Where("mid = ? AND word = ?", mid, word).Delete(&BlockWords{}).Error
}

func CheckBlockWordExist(mid string, word string) bool {
	var dummy BlockWords
	return !db.Model(&BlockWords{}).Where("mid = ? AND word = ?", mid, word).First(&dummy).RecordNotFound()
}

func UpdateBlockWord(mid string, word string, handle string) error {
	return db.Model(&BlockWords{}).Where("mid = ? AND word = ?", mid, word).Update("handle", handle).Error
}

type BlockWordsInfo struct {
	Mid     string
	Name    string
	Avatar  string
	Shield  string
	Handle  string
	Visible bool
}

func GetAllBlockWordsInfo() ([]BlockWordsInfo, error) {
	var infos []BlockWordsInfo
	result := db.Raw("select u.mid,u.name,u.avatar,b.word as shield,b.handle,b.visible from user u left join blocks_words b on u.mid=b.mid").Scan(&infos)
	return infos, result.Error
}

func SetWordVisibility(mid string, word string, visible bool) error {
	return db.Model(&BlockWords{}).Where("mid = ? AND word = ?", mid, word).Update("visible", visible).Error
}
