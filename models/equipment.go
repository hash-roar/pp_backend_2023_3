package models

import (
	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Mid      string
	Hostname string
	Version  string
	Jct      string
}

func GetEquipmentByHostnameMid(mid string, hostname string) (*Equipment, error) {
	var equipment Equipment
	result := db.Where(&Equipment{
		Hostname: hostname,
		Mid:      mid,
	}).First(&equipment)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		} else {
			return nil, nil
		}
	}
	return &equipment, nil

}

func CreateEquipment(e *Equipment) error {
	return db.Create(e).Error
}

func UpdateEquipmentJctVersion(hostname string, newJct string, version string) error {
	return db.Model(&Equipment{}).Where("hostname = ?", hostname).Updates(&Equipment{
		Version: version,
		Jct:     newJct,
	}).Error
}

func UpdateEquipment(where *Equipment, new *Equipment) error {
	return db.Model(where).Updates(new).Error
}

func GetAllEquipmentsByMid(mid string) ([]Equipment, error) {
	var equipments []Equipment
	result := db.Where("mid = ?", mid).Find(&equipments)
	return equipments, result.Error
}
